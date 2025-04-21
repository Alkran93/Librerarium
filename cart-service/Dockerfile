# Utiliza una imagen base de Go
FROM golang:1.24.3-bookworm


# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de dependencias y descarga las dependencias
COPY go.mod ./
RUN go mod download

# Copia el código fuente al contenedor
COPY . .

# Compila la aplicación Go
RUN go build -o main .

# Expone el puerto en el que la aplicación escuchará
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["./main"]
