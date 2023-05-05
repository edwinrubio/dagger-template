package main

import (
    "context"
    "fmt"
    "os"

    "dagger.io/dagger"
)

func main() {
    if err := build(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func buildd(ctx context.Context) error {
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
    src := client.Host().Directory(".")

    // create empty directory to put build outputs
    outputs := client.Directory()

    // get `golang` image
    golang := client.Container().From("golang:latest")

    // mount cloned repository into `golang` image
    golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

    for _, goos := range oses {
        for _, goarch := range arches {
            // creacion del directorio para la compilacion
            path := fmt.Sprintf("build/%s/%s/", goos, goarch)

            // setear GOARCH y GOOS para el entorno de compilacion
            build := golang.WithEnvVariable("GOOS", goos)
            build = build.WithEnvVariable("GOARCH", goarch)

            // compilar la aplicacion
            build = build.WithExec([]string{"go", "build", "-o", path})

            // obtener referencia al directorio de salida del contenedor
            outputs = outputs.WithDirectory(path, build.Directory(path))
        }
    }
    // write build artifacts to host
    _, err = outputs.Export(ctx, ".")
    if err != nil {
        return err
    }

    return nil
}