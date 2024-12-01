FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o loadtester main.go


FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/loadtester /usr/local/bin/loadtester
ENTRYPOINT ["loadtester"]
