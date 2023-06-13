FROM golang:1.18-alpine

WORKDIR /app/gp-footballer-svc
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]