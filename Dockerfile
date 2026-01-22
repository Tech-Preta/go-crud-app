# Etapa 1: Build
FROM golang:1.23-alpine3.20 AS builder

# Instalar gcc
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copie o restante dos arquivos do projeto
COPY . .

# Compile o aplicativo
RUN go build -o main .

# Etapa 2: Runtime
FROM alpine:3.23.2

WORKDIR /app

# Instalar curl para o healthcheck
RUN apk add --no-cache curl

# Copie o binário compilado da etapa de build
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Exponha a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]