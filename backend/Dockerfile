# ---------- BUILD STAGE ----------
FROM golang:1.25.4-alpine AS build

WORKDIR /app

# Copiamos primero dependencias del módulo
COPY go.mod go.sum ./
RUN go mod download

# Ahora copiamos todo el backend (SOLAMENTE el backend)
COPY . .

# Construimos el binario del backend
RUN go build -o server main.go

# ---------- RUN STAGE ----------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copiar binario compilado
COPY --from=build /app/server .

# Copiar env del backend
COPY .env . 

EXPOSE 8080

CMD ["./server"]
