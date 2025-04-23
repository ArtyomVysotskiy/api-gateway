# API Gateway

This is a REST API Gateway for the auth-api gRPC service. It provides REST endpoints that map to the underlying gRPC service calls.

## API Endpoints

### Authentication

- `POST /api/v1/auth/register`
  - Request body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "user_id": "string"
    }
    ```

- `POST /api/v1/auth/login`
  - Request body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "access_token": "string",
      "refresh_token": "string"
    }
    ```

- `POST /api/v1/auth/refresh`
  - Request body:
    ```json
    {
      "refresh_token": "string"
    }
    ```
  - Response:
    ```json
    {
      "access_token": "string"
    }
    ```

- `POST /api/v1/auth/validate`
  - Request body:
    ```json
    {
      "access_token": "string"
    }
    ```
  - Response:
    ```json
    {
      "valid": boolean
    }
    ```

- `POST /api/v1/auth/logout`
  - Request body:
    ```json
    {
      "access_token": "string"
    }
    ```
  - Response: Empty response with 200 status code

## Running the Service

### Prerequisites

- Docker
- Docker Compose

### Running with Docker Compose

1. Make sure you're in the `api-gateway` directory
2. Run:
   ```bash
   docker-compose up --build
   ```

The API Gateway will be available at `http://localhost:8080`

### Development

To run the service locally for development:

1. Install Go 1.21 or later
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the service:
   ```bash
   go run main.go handlers.go
   ```

## Environment Variables

- `AUTH_API_HOST`: The host and port of the auth-api gRPC service (default: auth-api:50051) 