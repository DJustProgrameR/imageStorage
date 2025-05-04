# imageStorage

### Генерация кода клиента
```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/ascii-pet.json -g typescript-axios -o /local/frontend/src/client --additional-properties=supportsES6=true
```

### Генерация кода сервера
```
Создаём конфиг
backend/server_generation.yaml

Подтягиваем нужную либу
go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen

запускаем go генерацию в файле
backend/generate.go
```

### Запуск линтера
```
в ./backend :

golangci-lint run ./...
```

### Запуск приложения
```
docker-compose up  --build
```

