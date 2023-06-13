FROM golang:1.18-alpine

WORKDIR /app/gateway-main
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]