FROM golang:1.23 AS builder

WORKDIR /app

# Instala dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código
COPY . .

# Compila o binário
RUN go build -o main cmd/main.go

# Executa o binário
CMD ["./main"]
