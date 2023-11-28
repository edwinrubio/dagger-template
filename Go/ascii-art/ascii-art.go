package main

import (
    "context"
    "fmt"
    "os"
    "dagger.io/dagger"
)

func main() {
    if err := sonarScan(context.Background()); err != nil {
        fmt.Println(err)
    }
    
    if err := build(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Building with Dagger")

    // define build matrix
    oses := []string{"linux", "darwin"}
    arches := []string{"amd64", "arm64"}

    // initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()

    // get reference to the local project
    src := client.Host().Directory("ascii-art-cmd")

    // create empty directory to put build outputs
    outputs := client.Directory()

    // get `golang` image
    golang := client.Container().From("golang:latest")

    // mount cloned repository into `golang` image
    golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

    for _, goos := range oses {
        for _, goarch := range arches {
            // create a directory for each os and arch
            path := fmt.Sprintf("build/%s/%s/", goos, goarch)

            // set GOARCH and GOOS in the build environment
            build := golang.WithEnvVariable("GOOS", goos)
            build = build.WithEnvVariable("GOARCH", goarch)

            // build application
            build = build.WithExec([]string{"go", "build", "-o", path})

            // get reference to build output directory in container
            outputs = outputs.WithDirectory(path, build.Directory(path))
        }
    }
    
    // Guardando los artefactos resultado de la compilacion
    _, err = outputs.Export(ctx, "./ascii-art-cmd")

    //Especificando el directorio de trabajo
	src = client.Host().Directory("ascii-art-cmd")

	if err != nil {
		fmt.Printf("Error obteniendo la refenrencia del directorio host: %s", err)
		os.Exit(1)
	}

    //Definiendo la imagen de trabajo
	golang = client.Container().From("golang:latest")

	cn, err := client.Container().
		Build(src).
		Publish(ctx, "allfait/ascii-art-cmd:latest")
    
	if err != nil {
		fmt.Printf("Error creando y subiendo el contenedor: %s", err)
		os.Exit(1)
    }

    fmt.Print("Contenedor creado y pusheado: %s", cn)

    if err != nil {
        return err
    }

    return nil
}

func sonarScan(ctx context.Context) error {
    fmt.Println("Scanning with Sonar")

    // initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()
    // get reference to the local project
    src := client.Host().Directory("ascii-art-cmd")

    // create empty directory to put build outputs
    outputs := client.Directory()

    //especificamos la imagen que vamos a usar
    sonar := client.Container().From("sonarsource/sonar-scanner-cli:4.8")

	sonar = sonar.WithDirectory("ascii-art-cmd", src).WithWorkdir("ascii-art-cmd")

    //Ejecutamos el cliente de sonarqube
    sonar = sonar.WithExec([]string{"sonar-scanner", "-X"})

    path := fmt.Sprintf(".")


    // get reference to build output directory in container
    outputs = outputs.WithDirectory(path, sonar.Directory(path))

    // write build artifacts to host
    _, err = outputs.Export(ctx, ".")
    if err != nil {
        return err
    }
    
    return nil

}