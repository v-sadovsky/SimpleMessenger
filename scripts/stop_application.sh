#!/usr/bin/env bash

# This script stops containerized application in the background

echo "==> Stopping SimpleMessenger application..."
docker compose -f ./deploy/docker-compose.yml down
echo "Done!"