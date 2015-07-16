#! /bin/bash

# fix inconsistent docker env var name
export MINIWAVEME_DB_URL=${DATABASE_1_PORT_27017_TCP_ADDR}
export MINIWAVEME_DB_PORT=${DATABASE_1_PORT_27017_TCP_PORT}

# Run the api
go run main.go
