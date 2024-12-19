# go-es-patterns 🧪

Event Sourcing and CQRS implementation patterns in Go. This project serves as a technical laboratory for experimenting with various event sourcing patterns, CQRS implementations, and performance optimizations.

## 🎯 Overview

This project focuses on:
- 📊 Event Sourcing pattern implementations
- 🔄 CQRS (Command Query Responsibility Segregation)
- 🔒 Various concurrency control mechanisms
- ⚡ Performance optimization techniques
- 💾 Event store implementations
- 📈 Projection strategies

## 🛠 Technical Stack

- ⚡ Go 1.21+
- 🐘 PostgreSQL (Event Store)
- 🐳 Docker & Docker Compose
- 🧪 Testing Tools
    - Go Testing Framework
    - Testcontainers
    - GoMock

## 📁 Project Structure

```
├── cmd/
│   └── api/                    # 애플리케이션 진입점
│       └── main.go
│
├── domain/                     # 모든 인터페이스와 도메인 엔티티
│   ├── entity/                # 도메인 엔티티
│   │   ├── account.go         # Account 애그리게잇
│   │   └── event.go           # 도메인 이벤트 정의
│   │
│   ├── repository/            # 저장소 인터페이스
│   │   ├── event_store.go     # 이벤트 저장소 인터페이스
│   │   └── read_store.go      # 읽기 모델 저장소 인터페이스
│   │
│   ├── service/               # 도메인 서비스 인터페이스
│   │   ├── event_publisher.go
│   │   └── event_handler.go
│   │
│   └── usecase/               # 유스케이스 인터페이스
│       ├── command/
│       │   ├── create_account.go
│       │   └── deposit_money.go
│       └── query/
│           ├── get_balance.go
│           └── get_history.go
│
├── router/                    # 라우팅 설정
│   └── router.go
│
├── controller/                # HTTP 요청 처리
│   ├── account_controller.go
│   └── dto/
│
├── usecase/                  # 유스케이스 구현체
│   ├── command/
│   └── query/
│
└── repository/               # 리포지토리 구현체
    └── postgres/
```

## 🔥 Implementation Features

### 📊 Event Sourcing
- Multiple event store implementations
- Event versioning
- Schema evolution
- Snapshots
- Optimistic concurrency control

### 🔄 CQRS
- Command handling
- Event handling
- Read model projections
- Asynchronous processing

### ⚡ Performance
- Snapshot strategies
- Caching mechanisms
- Batch processing
- Concurrent event handling

## 🚀 Getting Started

### Prerequisites
```bash
go 1.21+
docker
docker-compose
make
```

### Development Setup
```bash
# Clone repository
git clone https://github.com/username/go-es-lab.git
cd go-es-lab

# Start infrastructure
make docker-up

# Run tests
make test
```

## 💡 Running Examples

Basic account example demonstrating event sourcing:
```bash
make run-account-example
```

## 👨‍💻 Development

### Running Tests
```bash
# Unit tests
make test

# Integration tests
make integration-test

# Benchmark tests
make bench
```

### Adding New Patterns

1. Create a new directory under `internal/examples`
2. Implement the pattern
3. Add tests
4. Add documentation
5. Add benchmarks if applicable

## 🎯 Project Goals

- 🧪 Experiment with different event sourcing patterns
- 📊 Compare performance characteristics
- 📚 Provide reference implementations
- 📝 Document practical concerns and solutions
- 🤝 Share learnings about event sourcing in Go

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

### Development Process
1. 🔱 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'feat: Add some amazing feature'`)
4. 🚀 Push to the branch (`git push origin feature/amazing-feature`)
5. 🎉 Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.