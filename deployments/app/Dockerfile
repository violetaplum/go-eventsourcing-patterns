FROM golang:1.23.2-alpine

WORKDIR /app

# Kafka 클라이언트 라이브러리 빌드에 필요한 의존성 추가
RUN apk add --no-cache \
    git \
    librdkafka-dev \
    pkgconfig \
    build-base \
    bash

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]