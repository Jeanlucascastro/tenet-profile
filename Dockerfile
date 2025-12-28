# -------- BUILD STAGE --------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# dependências
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd

# -------- RUNTIME STAGE --------
FROM alpine:3.20

WORKDIR /app

# certificados (https / postgres)
RUN apk add --no-cache ca-certificates

# copia binário
COPY --from=builder /app/app .

# copia o .env (opcional)
COPY cmd/.env .env

EXPOSE 8082

CMD ["./app"]
