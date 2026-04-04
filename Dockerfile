FROM golang:1.25

# Define diretório de trabalho
WORKDIR /app

# Copia arquivos de dependência primeiro (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do projeto
COPY . .

# Expõe a porta
EXPOSE 8000

# Build da aplicação
RUN go build -o main ./cmd/api

# Executa
CMD ["./main"]