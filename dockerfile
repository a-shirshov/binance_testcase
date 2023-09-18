FROM golang:1.19-alpine as server_build
RUN apk add --no-cache \
    make \
    libwebp-dev \
    gcc \
    musl-dev \
    ca-certificates
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make binance_test
EXPOSE 8080
WORKDIR /app/bin/api
CMD ["./binance"]