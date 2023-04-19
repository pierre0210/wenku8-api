FROM golang:1.20.3-alpine AS builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./wenku8 ./cmd/wenku8-api/main.go

FROM alpine:latest

WORKDIR /
COPY --from=builder /app/wenku8 wenku8
EXPOSE 5000
CMD [ "./wenku8" ]