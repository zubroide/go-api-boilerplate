# Go API Boilerplate


## Features

- [x] Framework for API: Gin
- [x] Package manager: go mod
- [x] DI: Based on service container
- [x] Layers: Controller->Service->Entity
- [x] Routes: Gin
- [x] Process controller results and convert them into JSON/XML according to request headers
- [x] Logger: logrus
- [x] Environment variables, config: Viper
- [x] ORM: GORM
- [x] Migrations: goose
- [x] Data seeders
- [x] Console commands: Cobra
- [x] Unit tests with overriding of services in DI (`go test`)
- [x] Code coverage by tests (`go tool cover`)
- [x] Logger integration with Sentry
- [x] Setup alerting for unhandled errors
- [x] Swagger
- [x] Docker compose
- [x] Makefile
- [x] Development: hot reload code


## Folders structure

- `command/`: Console commands.
- `controller/`: Controllers for web requests processing.
- `db/`: Migrations and seeders.
- `dic/`: Dependency Injection Container.
- `doc/`: Swagger documentation.
- `docker/`: Docker containers description.
- `install/`: Scripts for environment preparing.
- `logger/`: Logger and client for Sentry.
- `model/`: Business logic.
- `model/db/`: DB connection.
- `model/entity/`: GORM entities.
- `model/service/`: Business logic.
- `route/`: Web requests routes.
- `vendor/`: Packages using in application.
- `base.env`: Base environment variables.
- `.env`: Environment variables for current environment.


## How to use (Docker)


```bash
docker-compose up --build
```

Check 
- http://localhost:8080/users
- http://localhost:8080/doc/swagger/index.html


## How to use (without Docker)

### Install necessary packages

```bash
./install/install.sh
```


### Create and edit config


```bash
cp .env.template .env
mcedit .env
```


### Get vendor packages

```bash
go mod vendor
```


### Run migrations

Create database `go-api-boilerplate`.

And run migrations:

```bash
make migrate
```


### Run application

Check available commands

```bash
make
```

Run http server

```bash
make server
```

Or:

```bash
go run main.go server --port=8081
```

Check http://localhost:8081


### Run tests

Run all tests:

```bash
go test ./... -v -coverpkg=./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Run test for one package:

```bash
go test go-api-boilerplate/test/unit -v -coverpkg=./... -coverprofile=coverage.out
```

Run one test:

```bash
go test test/unit/user_service_test.go -v -coverpkg=./... -coverprofile=coverage.out
```

Using make:

```bash
make test
```


### Generate Swagger documentation

Generate swagger.json:

```bash
make swagger
```

Documentation must be available at url http://localhost:8081/doc/swagger/index.html


## Requirements
  - Go 1.23+
