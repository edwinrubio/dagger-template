# Imagen base de Go
FROM golang:1.20.3-alpine

# Establecemos el directorio de trabajo en /app
WORKDIR /app

# Copiamos el contenido del directorio actual al directorio /app
COPY /build/linux/amd64 .
COPY imagen.jpeg .

# Exponemos el puerto 8080
EXPOSE 8085

# Ejecutamos la aplicación
CMD ["./api-ascii-art"]