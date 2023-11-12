#!/bin/bash

echo "Building the server..."
go build -o journal-api ./cmd/server

echo "Starting the server..."
./journal-api