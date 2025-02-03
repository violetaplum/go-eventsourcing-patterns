# deployments/app/Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# 필요한 시스템 패키지 설치
RUN apk add --no-cache git

# 소스 코드 복사
COPY . .

# 의존성 다운로드
RUN go mod download

# 애플리케이션 빌드
RUN go build -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]