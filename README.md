# golang-chatapp

chat-app-backend/
│
├── api-gateway/              # Kong / Envoy configuration
│   └── kong.yaml
│
├── services/
│   ├── user-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── models/
│   │   │   ├── repository/
│   │   │   └── services/
│   │   ├── proto/user.proto
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── chat-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   ├── models/
│   │   │   ├── repository/
│   │   │   └── services/
│   │   ├── proto/chat.proto
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── video-service/
│   │   ├── cmd/main.go
│   │   ├── internal/
│   │   │   ├── signaling/
│   │   │   └── webrtc/
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   └── notification-service/
│       ├── cmd/main.go
│       ├── internal/
│       │   ├── websocket/
│       │   └── push/
│       ├── Dockerfile
│       └── go.mod
│
├── pkg/                      # Shared utils: auth, logger, middleware
│   ├── auth/
│   ├── logger/
│   └── config/
│
├── deploy/
│   ├── k8s/
│   │   ├── namespace.yaml
│   │   ├── user-service.yaml
│   │   ├── chat-service.yaml
│   │   ├── video-service.yaml
│   │   ├── notification-service.yaml
│   │   ├── postgres.yaml
│   │   ├── mongo.yaml
│   │   ├── redis.yaml
│   │   └── ingress.yaml
│   └── docker-compose.yml
│
└── Makefile
