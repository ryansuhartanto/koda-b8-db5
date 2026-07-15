#!/bin/sh
docker-entrypoint.sh postgres &

env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /var/app/.env

until pg_isready -q -U "$POSTGRES_USER"; do sleep 1; done
exec /bin/app
