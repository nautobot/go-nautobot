#!/bin/bash
set -euxo pipefail

cleanup() {
    docker compose --project-name get_api -f docker-compose.yml down -v || true
}

trap cleanup EXIT

echo "Fetching openapi.yaml"

docker compose --project-name get_api -f docker-compose.yml --env-file "../local_dev.env" up --build --abort-on-container-exit --exit-code-from get-api

echo "openapi files created"
