FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/order-service ./cmd/order-service
WORKDIR /app/cmd/order-service
RUN go build -o /bin/order-service .

FROM alpine:latest
COPY --from=build /bin/order-service /usr/local/bin/order-service
EXPOSE 8082
ENTRYPOINT ["/usr/local/bin/order-service"]
