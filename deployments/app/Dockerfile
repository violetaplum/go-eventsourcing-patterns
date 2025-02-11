FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache \
    git \
    librdkafka-dev \
    pkgconfig \
    build-base \
    bash \
    gcc \
    musl-dev \
    linux-headers \
    libc-dev

COPY . .

RUN go mod download

# CGO_ENABLED=1을 명시적으로 설정하고, 빌드 시 추가 플래그 사용
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]