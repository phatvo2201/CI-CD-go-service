#!/bin/sh

set -e
echo "start source env file copy from secret aws"
source /app/app.env

echo "start run db migration in docker"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"