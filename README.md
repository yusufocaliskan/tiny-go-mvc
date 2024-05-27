# Introduction

A tiny-funny go framework created base on MVC structure.

## Instllation

Basically run the main.go
`go run main.go`

or use makefile
`make run`

The framework also support hot-reload using fresh, use the below codes.

`make hot`

To test units, run the docker-compose.dev. (Remember, testing units is a bit tricky, and you might encounter errors.)

`make test-all`

Or

```shell
docker-compose -f docker-compose.dev.yml exec gptv_backend go test -v ./...
```

Every time you run **make hot** the swagger files will be regenared

## Docker

Building docker with Development in hot reload using **fresh**

Building:

```shell
docker compose -f "docker-compose.dev.yml"  build
```

Run the dockers:

```shell
docker compose -f "docker-compose.dev.yml"  up
```

## Swagger

<http://localhost:8080/swagger/index.html>

## gRPC Serve
