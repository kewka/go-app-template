#!/bin/sh
echo "Waiting postgres"
./wait-for.sh ${POSTGRES_HOST}:${POSTGRES_PORT}
exec "$@"
