FROM docker:dind

RUN cd /usr/local \
    curl -L https://dl.dagger.io/dagger/install.sh | sh 

COPY --from=golang:1.20.3-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

ENV GO111MODULE="on"

ENTRYPOINT [ "/usr/local/bin/docker-entrypoint.sh" ]


