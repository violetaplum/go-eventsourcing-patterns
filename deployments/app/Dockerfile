FROM golang:1.22 AS builder

WORKDIR /app

# apt-get add 대신 apt-get install 사용, 캐시 정리 추가
RUN apt-get update && apt-get install -y \
    git \
    librdkafka-dev \
    pkg-config \
    build-essential \
    bash \
    gcc \
    libc-dev \
    && rm -rf /var/lib/apt/lists/*

COPY . .

RUN go mod download

# http router 테스트 실행
RUN go test -v ./...

# musl 태그 제거, 주석 처리된 줄 삭제
RUN CGO_ENABLED=1 GOOS=linux go build -o account ./cmd/account/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -o event ./cmd/event/main.go

FROM ubuntu:22.04 AS account-app
# apt-get add 대신 apt-get install 사용, 캐시 정리 추가
RUN apt-get update && apt-get install -y librdkafka-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/account /account
EXPOSE 8080

CMD ["/account"]

FROM ubuntu:22.04 AS event-app
# apt-get add 대신 apt-get install 사용, 캐시 정리 추가
RUN apt-get update && apt-get install -y librdkafka-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/event /event

CMD ["/event"]