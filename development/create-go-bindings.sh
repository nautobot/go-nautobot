#!/bin/bash
set -euxo pipefail

cleanup() {
    docker-compose --project-name go_nautobot -f docker-compose.yml down -v || true
}

trap cleanup EXIT

echo "Generating GO bindings"

docker-compose --project-name go_nautobot -f docker-compose.yml --env-file local_dev.env up --build --abort-on-container-exit --exit-code-from oapi

echo "Done"
