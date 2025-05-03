# imageStorage

### Client generation
```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/ascii-pet.json -g typescript-axios -o /local/frontend/src/client --additional-properties=supportsES6=true
```

### Server generation
```
Создаём конфиг
backend/server_generation.yaml

Подтягиваем нужную либу
go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen

запускаем go генерацию в файле
backend/generate.go
```