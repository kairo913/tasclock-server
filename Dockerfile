FROM golang:1.23.2 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM debian:bookworm-slim

COPY --from=builder /app/main ./app

EXPOSE 8080

CMD [ "./app" ]