# Go API Demo _(go-api-demo)_

This project demonstrates a basic REST API using Go, Fiber, MongoDB

## Install

```bash
go mod init github.com/nubrid/go-api-demo

go get -u github.com/gofiber/fiber/v2 github.com/go-playground/validator/v10

go get go.mongodb.org/mongo-driver/mongo

npx kill-port 3000 && go run cmd/main.go

.\skaffold.exe dev
```
