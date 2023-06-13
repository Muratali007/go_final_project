FROM golang:1.18-alpine

WORKDIR /app/go-club-svc
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]