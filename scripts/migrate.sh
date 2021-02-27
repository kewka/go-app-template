#!/bin/sh
CONTAINER_ID=$(docker-compose run -d --rm migrate sleep 300)
docker cp ./migrations $CONTAINER_ID:/migrations
docker cp ./scripts/wait-for.sh $CONTAINER_ID:/wait-for.sh
docker exec $CONTAINER_ID sh /wait-for.sh postgres:5432
docker exec $CONTAINER_ID migrate -path /migrations -database postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/$POSTGRES_DB?sslmode=disable $@
docker cp $CONTAINER_ID:/migrations/. ./migrations
docker rm --force $CONTAINER_ID