#!/bin/bash

set -e

create_database() {
  local database="$1"

  if [ -z "$database" ] || [ "$database" = "$POSTGRES_DB" ]; then
    return
  fi

  if ! psql \
    --username "$POSTGRES_USER" \
    --dbname "$POSTGRES_DB" \
    --tuples-only \
    --command "SELECT 1 FROM pg_database WHERE datname = '$database'" \
    | grep -q 1; then
    createdb --username "$POSTGRES_USER" "$database"
  fi
}

create_database "${DB_INVENTORY_NAME:-inventory_db}"
create_database "${DB_NOTIFICATION_NAME:-notification_db}"
create_database "${DB_PAYMENT_NAME:-payment_db}"
