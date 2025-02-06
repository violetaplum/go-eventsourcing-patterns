# deployments/app/Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# 필요한 시스템 패키지 설치
RUN apk add --no-cache git curl openjdk11-jre && \
    curl -O https://downloads.apache.org/kafka/3.5.1/kafka_2.13-3.5.1.tgz && \
    tar -xzf kafka_2.13-3.5.1.tgz && \
    mv kafka_2.13-3.5.1 /usr/local/kafka && \
    rm kafka_2.13-3.5.1.tgz

# Kafka CLI 경로를 추가
ENV PATH="${PATH}:/usr/local/kafka/bin"

# 소스 코드 복사
COPY . .

# 의존성 다운로드
RUN go mod download

# 애플리케이션 빌드
RUN go build -o main ./cmd/api

#Kafka 초기화 스크립트 복사
COPY deployments/app/kafka-init.sh /app/kafka-init.sh
RUN chmod +x /app/kafka-init.sh

EXPOSE 8080

CMD ["/bin/sh", "-c", "/app/kafka-init.sh && ./main"]