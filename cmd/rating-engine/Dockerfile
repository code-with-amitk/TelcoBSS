FROM debian:bookworm-slim AS build
RUN apt-get update && apt-get install -y cmake build-essential pkg-config libssl-dev librdkafka-dev libcouchbase-dev
WORKDIR /app
COPY ./ ./
RUN cmake -S . -B build
RUN cmake --build build --config Release

FROM debian:bookworm-slim
COPY --from=build /app/build/rating-engine /usr/local/bin/rating-engine
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/rating-engine"]
