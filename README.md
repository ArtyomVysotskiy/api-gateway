# API Gateway

Это REST API Gateway для gRPC сервисов. Он предоставляет REST эндпоинты, которые в свою очередь вызывают gRPC сервисы.

## API Эндпоинты

### auth-api

- `POST /api/v1/auth/register`
  - Тело запроса:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Ответ:
    ```json
    {
      "user_id": "string"
    }
    ```

- `POST /api/v1/auth/login`
  - Тело запроса:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Ответ:
    ```json
    {
      "access_token": "string",
      "refresh_token": "string"
    }
    ```

- `POST /api/v1/auth/refresh`
  - Тело запроса:
    ```json
    {
      "refresh_token": "string"
    }
    ```
  - Ответ:
    ```json
    {
      "access_token": "string"
    }
    ```

- `POST /api/v1/auth/validate`
  - Тело запроса:
    ```json
    {
      "access_token": "string"
    }
    ```
  - Ответ:
    ```json
    {
      "valid": boolean
    }
    ```

- `POST /api/v1/auth/logout`
  - Тело запроса:
    ```json
    {
      "access_token": "string"
    }
    ```
  - Ответ: Пустой ответ со статусом 200

## Запуск сервиса

### Предварительные требования

- Go 1.23.4 или выше

### Запуск с помощью Docker Compose

1. Убедитесь, что вы находитесь в директории `api-gateway`

2. Установите зависимости:
   ```bash
   go mod download
   ```
3. Запустите сервис:
   ```bash
   make run
   ```

API Gateway будет доступен по адресу `http://localhost:8080`
