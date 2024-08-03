#!/usr/bin/env bash

# This script starts containerized application in the background

echo "==> Starting SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml up -d
echo "SimpleMessenger has been started!"