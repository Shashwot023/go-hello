FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
