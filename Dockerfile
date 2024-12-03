FROM golang:1.20 AS builder
WORKDIR /app
RUN go mod download
COPY . .
RUN go build -o loadtester main.go


FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/loadtester /usr/local/bin/loadtester
ENTRYPOINT ["loadtester"]
