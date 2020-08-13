# Go API Boilerplate


## Features

- [x] Framework for API: Gin
- [x] Package manager: go mod
- [x] DI: Based on service container
- [x] Layers: Controller->Service->Repository->Entity
- [x] Routes: Gin
- [x] Process controller results and convert them into JSON/XML according to request headers
- [x] Logger: logrus
- [x] Environment variables, config: Viper
- [x] ORM: GORM
- [x] Migrations: gorm-goose
- [x] Base CRUD service
- [x] Base CRUD repository
- [x] Base CRUD controller
- [x] Request validation (Gin)
- [x] Console commands: Cobra
- [x] Unit tests with overriding of services in DI (`go test`)
- [x] Code coverage by tests (`go tool cover`)
- [x] Logger integration with Sentry: logrus_sentry
- [x] Setup alerting for unhandled errors
- [x] Swagger
- [x] Docker compose


## Folders structure

- `command/`: Console commands.
- `controller/`: Controllers for web requests processing.
- `db/`: Migrations.
- `dic/`: Dependency Injection Container.
- `doc/`: Swagger documentation.
- `docker/`: Docker containers description.
- `install/`: Scripts for environment preparing.
- `logger/`: Logger and client for Sentry.
- `model/`: Business logic.
- `model/db/`: DB connection.
- `model/entity/`: GORM entities.
- `model/repository/`: Repositories for access to storage.
- `model/service/`: Business logic.
- `route/`: Web requests routes.
- `test/`: Unit tests.
- `vendor/`: Packages used in application.
- `.env`: Environment variables for current environment.
- `base.env`: Base environment variables.


## How to use (Docker)


```bash
docker-compose up --build
```

Check http://localhost:8080


## How to use (without Docker)


### Prepare environment for Go projects if you do not done it early

```bash
sudo apt update
sudo apt upgrade
# See last version here: https://golang.org/dl/
wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
sudo tar -xvf go1.12.5.linux-amd64.tar.gz
sudo mv go /usr/local
sudo mcedit /etc/profile
```

And add last line:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Update environment variables:

```bash
source /etc/profile
```

Check Go version:

```bash
go version
```

Now create folder for Go projects:

```bash
mkdir ~/go
cd ~/go
touch init.sh
mcedit init.sh
```

Paste next code into this file:

```bash
#!/bin/bash

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Execute file:

```bash
chmod +x init.sh
source init.sh
```


### Clone repo

```bash
git clone git@github.com:zubroide/go-api-boilerplate.git
cd go-api-boilerplate
```


### Install necessary packages

```bash
./install/install.sh
```


### Create and edit config


```bash
cp .env.example .env
mcedit .env
```


### Download vendor packages

```bash
go mod download
```


### Run migrations

Create database `go-api-boilerplate`.

And run migrations:

```bash
make migrate
```


### Run application

```bash
make server
```

Or:

```bash
go run main.go server --port=8080
```

Check http://localhost:8080


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

Documentation must be available at url http://localhost:8080/swagger/index.html


## Requirements
  - Go 1.12+
