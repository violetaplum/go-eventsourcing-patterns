#!/bin/sh

# Kafka 브로커 준비 대기
echo "Waiting for Kafka to be ready..."
while ! nc -z kafka 9092; do
  sleep 1
done
echo "Kafka is ready!"

# Kafka 토픽 생성
echo "Creating Kafka topic: account-events..."
kafka-topics.sh --create --if-not-exists \
  --bootstrap-server kafka:9092 \
  --replication-factor 1 \
  --partitions 1 \
  --topic account-events
echo "Kafka topic created successfully!"

# 애플리케이션 실행
echo "Starting the application..."
exec "$@"