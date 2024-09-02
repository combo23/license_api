# license-api

License API is an open-source API designed to help manage licenses for your product. This project is built with Go.

## Getting Started

### Prerequisites
- go
- `.env` based on .env.example
- `settings.yaml` 

### Installation and Usage

### Using Docker
The recommended way to deploy License API is via Docker. To get started, simply run:
```bash
docker compose up
```

### Building from source

If you prefer to build the application from source, you can do so with the following command:
```bash
go build -o main.exe cmd/api/main.go
```
Note that when building from source, you must deploy your own instance of the database. While MongoDB is typically managed via Docker, you'll need to ensure your database is properly configured and running if you opt out of Docker.

## Database 

License API uses [MongoDB](https://www.mongodb.com/) to store and manage licenses. The licenses are stored in the following JSON format:

```json
{
    "license": "XXXXX-XXXXX-XXXXX-XXXXX-XXXXX",
    "username": "combo23",
    "createdAt": "1723018924",
    "updatedAt": "1723018924",
    "expiresAt": "1723018924",
    "hwid":"cdd31a3e-d3c7-45cc-b271-e3ce4d1e568e",
    "status":"active"
}
```

## Examples

Example implementations of this licensing system can be found in the examples directory. These examples demonstrate how to integrate the License API into your own product.

## API Documentation

Comprehensive documentation about the API endpoints is available in the [API_documention.md](API_documentation.md) file

## MakeFile

The Makefile included with this project provides several commands to streamline common tasks:

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```