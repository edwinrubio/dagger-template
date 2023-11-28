package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura para representar un recurso (en este caso, un mensaje)
type Mensaje struct {
	Contenido string `json:"contenido"`
}

// Controlador para manejar las solicitudes HTTP en la ruta /mensaje
func mensajeHandler(w http.ResponseWriter, r *http.Request) {
	// Crear una instancia de la estructura Mensaje
	mensaje := Mensaje{Contenido: "Hola, este es un mensaje de la API REST en Go"}

	// Convertir la estructura a formato JSON
	mensajeJSON, err := json.Marshal(mensaje)
	if err != nil {
		http.Error(w, "Error al convertir el mensaje a JSON", http.StatusInternalServerError)
		return
	}

	// Establecer las cabeceras de respuesta
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON al cliente
	w.Write(mensajeJSON)
}

func main() {
	// Configurar el controlador para la ruta /mensaje
	http.HandleFunc("/mensaje", mensajeHandler)

	// Iniciar el servidor en el puerto 8080
	puerto := 8080
	fmt.Printf("Servidor escuchando en el puerto %d...\n", puerto)
	err := http.ListenAndServe(fmt.Sprintf(":%d", puerto), nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
