include .env
export

env = local

.env:
	@ln -sf ./configs/.env.$(env) .env
	
deps:
	@docker-compose up -d postgres
	@until docker-compose exec -T postgres pg_isready > /dev/null 2>&1; do echo "Waiting postgres"; done

psql:
	@docker-compose exec postgres psql -U ${POSTGRES_USER} ${POSTGRES_DB}

swagger:
	@swag init -g pkg/service/http_handler.go -o api/service

migrate-up:
	@./scripts/migrate.sh up

unit-test:
	@go test -short -v -race ./...

test-deps: deps
test-deps:
	@docker-compose exec -T postgres psql -U ${POSTGRES_USER} -c "DROP DATABASE IF EXISTS ${POSTGRES_DB}"
	@docker-compose exec -T postgres psql -U ${POSTGRES_USER} -c "CREATE DATABASE ${POSTGRES_DB}"
	@./scripts/migrate.sh up

test: export POSTGRES_DB=app_test
test: test-deps
test:
	@go test -v -cover -race ./...