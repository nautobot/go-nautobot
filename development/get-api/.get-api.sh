#!/bin/bash
set -euxo pipefail

echo "Getting openAPI spec"

BETA_TAG="beta"
ALPHA_TAG="alpha"
RC_TAG="rc"

# TODO: eventually we would like to generate for experimental Nautobot versions
if [[ "$NAUTOBOT_VER" == *"$BETA_TAG"* ]] || [[ "$NAUTOBOT_VER" == *"$ALPHA_TAG"* ]] || [[ "$NAUTOBOT_VER" == *"$RC_TAG"* ]]; then
  echo "${NAUTOBOT_VER} is not an official Nautobot version, no new bindings are generated."
  exit 0
fi

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}

NAUTOBOT_TOKEN=0123456789abcdef0123456789abcdef01234567
wget --tries=5 --header="Authorization: Token ${NAUTOBOT_TOKEN}" \
     -O /client/api/openapi.yaml \
     "http://nautobot:8080/api/swagger.yaml?api_version=${MAJOR_MINOR_VER}" || {
  echo "Failed to download openapi.yaml"
  exit 1
}

cp /client/api/openapi.yaml /client/api/openapi-original.yaml

echo "openAPI files created"
