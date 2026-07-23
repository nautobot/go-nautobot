#!/bin/bash
set -euxo pipefail

echo "Preparing generation"

# Remove generated files
for F in $(cat /client/.openapi-generator/FILES) ; do
    rm -f /client/"${F}"
done

cp /client/api/openapi-original.yaml /client/api/openapi.yaml

echo "Fixing spec file"

#Fix openapi spec file
/client/development/create-bindings/scripts/fix-spec.py

echo "Creating GO bindings"

#yaml file is too long
export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
openapi-generator-cli generate --config /client/development/create-bindings/oapi-config.yaml \
    --input-spec /client/api/openapi.yaml \
    --output /client/ \
    --inline-schema-options RESOLVE_INLINE_ENUMS=true \
    --http-user-agent go-nautobot/$(cat /client/api/nautobot_version)

rm /client/.travis.yml

echo "Adding missing imports"
/client/development/create-bindings/scripts/add-missing-imports.sh

echo "Copying READMEs"
mv /client/README.md /client/docs/README.md
sed -i 's|docs/||g' /client/docs/README.md
echo "docs/README.md" >> /client/.openapi-generator/FILES
cp /client/.README.md /client/README.md

echo "Starting Nautobot client tests..."

cd /client
go mod tidy
go test -v -gcflags="-e" ./...


echo "Nautobot client tests completed"
