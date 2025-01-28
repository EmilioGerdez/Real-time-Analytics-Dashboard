# Kafka Logger Example

A production-ready example demonstrating distributed logging using Kafka for message streaming, Redis for log storage, and Go for high-performance processing. Designed to showcase modern event-driven architecture patterns.

## Features

- 📨 **Kafka Integration**: Reliable log streaming with Sarama producer/consumer
- 🗄️ **Redis Storage**: Persistent log storage 
- 🐳 **Dockerized Infrastructure**: Single-command environment setup
- 🔄 **Scalability**: Horizontal scaling patterns for consumer groups

## Technologies

- **Apache Kafka**: Distributed event streaming platform
- **Redis**: In-memory data structure store
- **Golang**: High-performance producer/consumer implementation
- **Docker**: Containerized services orchestration

## Prerequisites

- Docker 20.10+
- Go 1.20+
- make (optional)
- Kafka CLI tools (optional)

## Quick Start

### 1. Clone Repository
```bash
git clone https://github.com/EmilioGerdez/Kafka-Logger-Example.git
cd Kafka-Logger-Example
```

### 2.
```bash 
docker-compose up -d --build
```

### 3.
```bash
go run cmd/log-reader/main.go
```

### 4.
```bash
go run cmd/example-app/main.go
```

License
----
MIT

> [!NOTE]
> This is a simplified example for demonstration purposes. Production implementations should consider additional error handling, monitoring, and security measures.
