#!/bin/sh
set -euxo pipefail

echo "Starting generation of Go Nautobot Bindings"

VERSION_FILE="tag.md"
CURRENT_VERSION=$(head -n 1 $VERSION_FILE)
CURRENT_MAJOR_MINOR_VER=${CURRENT_VERSION%.*}

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}

wget --tries=5  http://nautobot:8080/api/swagger.yaml?api_version=${MAJOR_MINOR_VER} -O swagger.yaml

oapi-codegen --config oapi-config.yml swagger.yaml

if [ "$CURRENT_MAJOR_MINOR_VER" = "$MAJOR_MINOR_VER" ]; then
    # Get the Patch version string
    NEW_PATCH_VERSION=$(echo $CURRENT_VERSION | awk -F '.' '{ print $3;}')
    # Remove suffixes in version
    NEW_PATCH_VERSION=${NEW_PATCH_VERSION%-*}
    # Increment Patch version string with 1
    NEW_PATCH_VERSION=$((${NEW_PATCH_VERSION} + 1))
    NEW_TAG=${CURRENT_MAJOR_MINOR_VER}.$NEW_PATCH_VERSION
else
    NEW_TAG=${MAJOR_MINOR_VER}.0
fi

# TODO: remove beta when it's in production
FINAL_NEW_TAG=${NEW_TAG}-beta

echo $FINAL_NEW_TAG > tag.md

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
