#!/bin/sh
set -euxo pipefail

echo "Starting generation of Go Nautobot Bindings"

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}

wget --tries=5  http://nautobot:8080/api/swagger.yaml?api_version=${MAJOR_MINOR_VER} -O swagger.yaml

oapi-codegen --config oapi-config.yml swagger.yaml

echo $NAUTOBOT_VER > tag.md

cp tag.md /client
cp *.go /client

echo "Go Nautobot Bindings generated"

echo "Starting Nautobot client tests..."

export NAUTOBOT_URL=http://nautobot:8080/api/
export NAUTOBOT_TOKEN=0123456789abcdef0123456789abcdef01234567

cd /client
go mod tidy
go test -v

echo "Nautobot client tests completed"
