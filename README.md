# echo-first-app

A simple REST API using Echo framework in Go for managing products.

## Run locally

```bash
go run cmd/server/main.go
```

## Features

### Create Product

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "test", "price": 123}' http://localhost:3000/products
```
