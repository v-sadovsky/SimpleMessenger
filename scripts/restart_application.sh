#!/usr/bin/env bash

# This script restarts containerized application

echo "==> Stopping SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml down
echo "==> Starting SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml up -d
echo "SimpleMessenger has been started!"