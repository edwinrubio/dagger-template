package docs

import (
    "dagger.io/dagger"
    "universe.dagger.io/bash"
    "universe.dagger.io/docker"
)

dagger.#Plan & {
    client: {
        filesystem: "./": read: {
            contents: dagger.#FS
            exclude: [
                "README.md",
                "eden.cue",
            ]
        }
        env: {
            SONAR_LOGIN:     dagger.#Secret
            GITHUB_HEAD_REF: GITHUB_HEAD_REF
        }
    }
    actions: {
        deps: {
            node:
                docker.#Build & {
                    steps: [
                        docker.#Pull & {
                            source: "index.docker.io/node"
                        },
                        docker.#Copy & {
                            contents: client.filesystem."./".read.contents
                            dest:     "./src"
                        },
                    ]
                }
            sonarscanner:
                docker.#Build & {
                    steps: [
                        docker.#Pull & {
                            source: "index.docker.io/sonarsource/sonar-scanner-cli"
                        },
                        docker.#Copy & {
                            contents: client.filesystem."./".read.contents
                            dest:     "/usr/src"
                        },
                    ]
                }
        }

        build: {
            bash.#Run & {
                workdir: "./src"
                input:   deps.node.output
                script: contents: #"""
                    npm ci
                    """#
            }
        }

        staticAnalysis: {
            lint:
                bash.#Run & {
                    workdir: "./src"
                    input:   build.output
                    script: contents: #"""
                        npx eslint --color .
                        """#
                }
            sonarscanner:
                docker.#Run & {
                    env: {
                        GITHUB_BRANCH_NAME: "main"
                        SONAR_LOGIN:        client.env.SONAR_LOGIN
                        SONAR_HOST_URL:     "http://192.168.0.6:9000"
                    }
                    workdir: "/usr/src"
                    input:   deps.sonarscanner.output
                }
        }

        test: {
            integrationTest: {
                workdir: "./src"
                docker.#Run & {
                    input: build.output
                    command: {
                        name: "/bin/bash"
                        args: ["-c", "npm run test:ci"]
                    }
                }
            }
            unitTest: {
                workdir: "./src"
                docker.#Run & {
                    input: build.output
                    command: {
                        name: "/bin/bash"
                        args: ["-c", "npm run test:unit"]
                    }
                }
            }
        }

        SCA: dependencyScanning: {
            workdir: "./src"
            docker.#Run & {
                input: build.output
                command: {
                    name: "/bin/bash"
                    args: ["-c", "npx audit-ci --high"]
                }
            }
        }

        #PythonBuild: docker.#Dockerfile & {
            dockerfile: contents: """
                    FROM node:16

                    # Create app directory
                    WORKDIR /usr/src/app

                    # Install app dependencies
                    # A wildcard is used to ensure both package.json AND package-lock.json are copied
                    # where available (npm@5+)
                    COPY package*.json ./

                    RUN npm ci --only=production
                    # If you are building your code for production
                    # RUN npm ci --only=production

                    # Bundle app source
                    COPY . .

                    EXPOSE 8080
                    CMD [ "node", "app.js" ]
                """
}
    }
}