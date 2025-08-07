# Chat App Backend 🗨️

A microservices-based backend system for a real-time chat application supporting messaging, video calls, and notifications, built using **Go**, **gRPC**, **WebSockets**, **WebRTC**, **Docker**, and **Kubernetes**.

---

## 📁 Project Structure

```bash
chat-app-backend/
│
├── api-gateway/               # Entry point for routing client requests to services
│   └── Dockerfile
│
├── go.mod                     # Root Go module file
├── main.go                    # Root main (can be used for gateway/bootstrap)
│
├── services/                  # Microservices
│   ├── user-service/          # Handles user registration, login, profile
│   ├── chat-service/          # Handles messaging and chat rooms
│   ├── video-service/         # WebRTC-based video call handling
│   └── notification-service/  # Real-time and push notifications
│
├── pkg/                       # Shared packages (auth, config, logger)
│
├── deploy/                    # Deployment configurations
│   ├── k8s/                   # Kubernetes YAMLs
│   └── docker-compose.yml     # Docker Compose for local dev/testing
│
└── Makefile                   # CLI automation tasks
```

---

## 🧩 Microservices Overview

| Service               | Description                                         | Tech Stack               |
|-----------------------|-----------------------------------------------------|--------------------------|
| **API Gateway**       | Routes external requests to internal services       | Go, gRPC, Docker         |
| **User Service**      | Authentication, user management                     | Go, gRPC, PostgreSQL     |
| **Chat Service**      | One-on-one / group chat, message persistence        | Go, gRPC, MongoDB        |
| **Video Service**     | Real-time video chat via WebRTC                     | Go, WebRTC               |
| **Notification Service** | Real-time alerts via WebSocket + push notifications | Go, WebSocket, Redis     |

---

## 🚀 Running Locally

### Prerequisites

- Docker
- Docker Compose

### Start All Services

```bash
docker-compose up --build
```

OR use the Makefile:

```bash
make up
```

### Tear Down

```bash
docker-compose down
```

---

## ☸️ Kubernetes Deployment

Ensure you have `kubectl` and a K8s cluster running (e.g., Minikube, kind, or cloud provider).

### Apply Resources

```bash
kubectl apply -f deploy/k8s/
```

---

## 📦 Dependencies

Each service is independently modularized with its own `go.mod`. Common dependencies include:

- **gRPC**
- **Protobuf**
- **GORM** (User Service)
- **MongoDB Driver** (Chat Service)
- **Redis Client** (Notification Service)
- **Pion WebRTC** (Video Service)

---

## 🔧 Shared Packages (pkg/)

- `auth/`: JWT token generation/validation
- `logger/`: Centralized logging
- `config/`: Shared configuration and environment loading

---

## 📜 Proto Files

Defined inside each service under `proto/`:

- `user.proto`
- `chat.proto`

These define the gRPC service contracts and messages.

---

## 🧪 Testing

Each service can be independently tested. For example, to run user-service tests:

```bash
cd services/user-service
go test ./...
```

---

## 📂 Docker Structure

All services have individual `Dockerfile`s. They're orchestrated together using:

- `docker-compose.yml` for local development
- Kubernetes YAMLs for production-grade orchestration

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add your message'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Create a Pull Request

---

## 📄 License

This project is licensed under the MIT License.
