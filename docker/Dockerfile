FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api ./cmd/api/main.go

FROM alpine:3.21

COPY --from=builder /bin/api /bin/api

COPY . /app

EXPOSE 8080

CMD ["./bin/api"]