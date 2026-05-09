FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/billing-service ./cmd/billing-service
WORKDIR /app/cmd/billing-service
RUN go build -o /bin/billing-service .

FROM alpine:latest
COPY --from=build /bin/billing-service /usr/local/bin/billing-service
EXPOSE 8083
ENTRYPOINT ["/usr/local/bin/billing-service"]
