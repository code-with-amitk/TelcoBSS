#!/usr/bin/env bash
set -euo pipefail

MYSQL_HOST=${MYSQL_HOST:-localhost}
MYSQL_PORT=${MYSQL_PORT:-3306}
MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-password}
MYSQL_DB=${MYSQL_DATABASE:-telco_bss}

export MYSQL_PWD="$MYSQL_PASSWORD"

mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u "$MYSQL_USER" <<SQL
CREATE DATABASE IF NOT EXISTS telco_bss;
USE telco_bss;
SOURCE $(pwd)/docs/database/schema/mysql_init.sql;
SQL

echo "MySQL database initialized."
