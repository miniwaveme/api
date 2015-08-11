#! /bin/bash

# fix inconsistent docker env var name
export MINIWAVEME_DB_URL=${DATABASE_1_PORT_27017_TCP_ADDR}
export MINIWAVEME_DB_PORT=${DATABASE_1_PORT_27017_TCP_PORT}

export MINIWAVEME_REDIS_URL=${REDIS_1_PORT_6379_TCP_ADDR}
export MINIWAVEME_REDIS_PORT=${REDIS_1_PORT_6379_TCP_PORT}
export MINIWAVEME_REDIS_NETWORK=${REDIS_1_PORT_6379_TCP_PROTO}

# Run the api
go run main.go
