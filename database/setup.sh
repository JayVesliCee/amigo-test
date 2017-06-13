#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER amigo;
    ALTER USER amigo WITH PASSWORD 'amigotmppwd';
    CREATE DATABASE amigodb;
    GRANT ALL PRIVILEGES ON DATABASE amigodb TO amigo;

EOSQL
