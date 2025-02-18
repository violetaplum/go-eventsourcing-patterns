FROM golang:1.22-alpine AS builder

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
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -o account ./cmd/account/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -o event ./cmd/event/main.go

FROM alpine:3.18 AS account-app
RUN apk add --no-cache librdkafka-dev
WORKDIR /app
COPY --from=builder /app/account /account
EXPOSE 8080

CMD ["/account"]

FROM alpine:3.18 AS event-app
RUN apk add --no-cache librdkafka-dev
WORKDIR /app
COPY --from=builder /app/event /event

CMD ["/event"]