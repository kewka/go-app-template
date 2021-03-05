# Go app template

## Requirements

- Go **1.16+**
- docker-compose **1.27.0+**

## Configuration

```sh
# Create .env file (default: local)
$ make -B .env
```

## Testing

```sh
# only unit tests
$ make unit-test

# all tests
$ make test
```

## Service HTTP API

[Swagger](api/service/swagger.yaml)

## Usage

```sh
$ ./app -help
Usage of ./app:
  -port string
        http server port (default "4000")
```

### Locally

```sh
$ make -B .env env=local
$ make deps
$ make migrate-up
# build web spa
$ cd web/ && yarn && yarn build && cd -
$ go run ./cmd/app
```

### Dockerizing

```sh
$ make -B .env env=docker
$ make migrate-up
$ docker-compose up app
```

## Database migrations

[migrate](https://github.com/golang-migrate/migrate) - CLI and Golang library.

### Up

```sh
$ ./scripts/env.sh ./scripts/migrate.sh up
```

### Down

```sh
$ ./scripts/env.sh ./scripts/migrate.sh down
```

### Create

```sh
$ ./scripts/env.sh ./scripts/migrate.sh create -dir /migrations -ext sql -seq create_new_table

Output:
migrations/XXXXXX_create_new_table.up.sql
migrations/XXXXXX_create_new_table.down.sql
```
