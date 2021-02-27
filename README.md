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

## HTTP API Server

[Swagger](pkg/httphandler/swagger/swagger.yaml)


### Usage
```sh
$ ./api-server -help
Usage of ./api-server:
  -port string
        server port (default "4000")
```

### Locally
```sh
# Create env file
$ make -B .env env=local
# Run dependencies (postgres, redis, etc...)
$ make deps
# Apply database migrations
$ ./scripts/env.sh ./scripts/migrate.sh up
# Run api-server
$ go run ./cmd/api-server
```

### Dockerizing
```sh
# Create env file
$ make -B .env env=docker
# Apply database migrations
$ ./scripts/env.sh ./scripts/migrate.sh up
# Up service
$ docker-compose up api-server
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