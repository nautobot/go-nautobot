#!/bin/sh
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

VERSION_FILE="/client/tag.md"
CURRENT_VERSION=$(head -n 1 $VERSION_FILE)
CURRENT_MAJOR_MINOR_VER=${CURRENT_VERSION%.*}

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}

NAUTOBOT_TOKEN=0123456789abcdef0123456789abcdef01234567
wget --tries=5 --header="Authorization: Token ${NAUTOBOT_TOKEN}" \
     -O swagger.yaml \
     "http://nautobot:8080/api/swagger.yaml?api_version=${MAJOR_MINOR_VER}" || {
  echo "Failed to download swagger.yaml"
  exit 1
}

echo "Creating GO bindings"

wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.58/swagger-codegen-cli-3.0.58.jar -O swagger-codegen-cli.jar

#yaml file is too long
export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
java -jar swagger-codegen-cli.jar generate -i swagger.yaml -l go -o /client -DpackageName=nautobot


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

echo "Starting Nautobot client tests..."

export NAUTOBOT_URL=http://nautobot:8080/api/
export NAUTOBOT_TOKEN=0123456789abcdef0123456789abcdef01234567

cd /client/test
go mod tidy
go test -v -gcflags="-e"

echo "Nautobot client tests completed"
