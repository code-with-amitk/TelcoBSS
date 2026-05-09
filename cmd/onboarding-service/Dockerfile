FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/onboarding-service ./cmd/onboarding-service
WORKDIR /app/cmd/onboarding-service
RUN go build -o /bin/onboarding-service .

FROM alpine:latest
COPY --from=build /bin/onboarding-service /usr/local/bin/onboarding-service
EXPOSE 8081
ENTRYPOINT ["/usr/local/bin/onboarding-service"]
