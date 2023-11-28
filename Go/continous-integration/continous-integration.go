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

    if err := sonarScan(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Construyendo con dagger")

    // Definiendo la matriz de compilacion
    oses := []string{"linux", "darwin"}
    arches := []string{"amd64", "arm64"}

    // Inicializando el cliente de dagger
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()

    // Obteniendo la referencia la proyecto local
    src := client.Host().Directory("./aplication/aplication.go")

    // Creando un directorio vacio para poner las salidas
    outputs := client.Directory()

    // obteniendo la ultima imagen de golang
    golang := client.Container().From("golang:latest")

    // montando el codigo en el contenedor de go
    golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

    for _, goos := range oses {
        for _, goarch := range arches {
            // creando un directorio para cada tipo de arquitectura
            path := fmt.Sprintf("build/%s/%s/", goos, goarch)

            // seteando GOARCH y GOOS en el ambiente de compilacion
            build := golang.WithEnvVariable("GOOS", goos)
            build = build.WithEnvVariable("GOARCH", goarch)

            // compilar la aplicacion
            build = build.WithExec([]string{"go", "build", "-o", path})

            // obteniendo la referencia de la compilacion de salida en el directorio del contenedor
            outputs = outputs.WithDirectory(path, build.Directory(path))
        }
    }
    
    // guardando los artefactos del compilacion en el host
    _, err = outputs.Export(ctx, ".")


    if err != nil {
        return err
    }

    return nil
}


func sonarScan(ctx context.Context) error {
    fmt.Println("Scanning with Sonar")

    // Inicilizando el cliente de dagger
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()

    src := client.Host().Directory(".")

    // Creando un directorio para poner las salidas
    outputs := client.Directory()

    // Obteniendo la imagen del cliente de sonnar scanner
    sonar := client.Container().From("sonarsource/sonar-scanner-cli:4.8")

    // montando el codigo en el contenedor de sonar
	sonar = sonar.WithDirectory(".", src).WithWorkdir(".")

    // creando un directorio para guardar lasa salidas
    path := fmt.Sprintf(".")

    //ejecutamos el agente de sonnar
	sonar = sonar.WithExec([]string{"sonar-scanner"})

    // obteniendo la referencia de la compilacion de salida en el directorio del contenedor
    outputs = outputs.WithDirectory(path, sonar.Directory(path))

    // guardando las salidas dentro del host
    _, err = outputs.Export(ctx, ".")
    if err != nil {
        return err
    }

    return nil
}