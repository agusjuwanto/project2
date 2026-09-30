# Gin REST API Template

Template REST API menggunakan Go + Gin.

## Struktur

```text
.
├── cmd/api/main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── response/
│   ├── route/
│   └── service/
├── .env.example
├── go.mod
└── README.md
```

## Menjalankan

```bash
cp .env.example .env
go mod tidy
go run ./cmd/api
```

Server:
`http://localhost:8080`

## Endpoint

- `GET /health`
- `GET /api/v1/users`
- `GET /api/v1/users/:id`
- `POST /api/v1/users`
- `PUT /api/v1/users/:id`
- `DELETE /api/v1/users/:id`

Contoh:

```bash
curl http://localhost:8080/api/v1/users

curl -X POST http://localhost:8080/api/v1/users   -H "Content-Type: application/json"   -d '{"name":"Budi","email":"budi@example.com"}'
```

Repository pada template ini masih in-memory agar mudah dipelajari. Untuk production, bagian repository dapat diganti PostgreSQL/MySQL menggunakan interface yang sama.
