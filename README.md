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
.
├── Makefile
├── README.md
├── application
│   ├── command
│   │   └── account_service.go
│   └── query
│       └── account_service.go
├── cmd
│   ├── account
│   │   └── main.go
│   └── event
│       └── main.go
├── deployments
│   ├── app
│   │   └── Dockerfile
│   ├── docker-compose.yml
│   └── postgres
│       └── init.sql
├── domain
│   ├── account.go
│   ├── config.go
│   ├── errors.go
│   ├── event.go
│   └── transaction.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── kafka
│   │   ├── consumer.go
│   │   ├── event_handler.go
│   │   └── producer.go
│   └── persistence
│       └── postgres
│           ├── account_store.go
│           ├── connection.go
│           └── event_store.go
├── interface
│   └── http
│       └── account_handler.go
└── tests
    └── api_test.http
```

## 🔥 Implementation Features

### 📊 Event Sourcing
- Multiple event store implementations
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