#!/usr/bin/env bash

# This script rebuild containerized application and starts it

echo "==> Stopping SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml down
echo "==> Building and starting SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml up --build -d
echo "SimpleMessenger has been built and started!"
