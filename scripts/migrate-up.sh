#!/bin/bash

readonly db_name="$1"
readonly migration_folder="$2"

migrate -database "postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${db_name}?sslmode=disable" \
  -path "./migrations/${migration_folder}" \
  up
