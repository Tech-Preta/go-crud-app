FROM golang:1.23-alpine3.20

# Instalar gcc
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copie o restante dos arquivos do projeto
COPY . .

# Compile o aplicativo
RUN go build -o main .

# Exponha a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]