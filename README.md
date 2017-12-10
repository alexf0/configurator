Загрузить Dep для vendor

```go get -u github.com/golang/dep/cmd/dep```

Установить все зависимости в корне приложения

```dep ensure```

Загрузить пакет для выполнения миграций

```go get -u -d github.com/mattes/migrate/cli github.com/lib/pq```
```go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli```

Применить все миграции вместе с seeds

```migrate -database postgres://postgres:password@localhost:5432?sslmode=disable -path ./migrations up```

Откатить последнюю миграцию с seeds

```migrate -database postgres://postgres:password@localhost:5432?sslmode=disable -path ./migrations down 1```

Проверить работу 

```curl -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/params -d '{"type": "Test.vpn","data":"Rabbit.log"}'```