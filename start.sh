#!/bin/sh

set -e
echo "start run db migration in docker"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"