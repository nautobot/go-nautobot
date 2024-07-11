#!/bin/bash
set -euxo pipefail

cleanup() {
    docker-compose --project-name go_nautobot -f docker-compose.yml down -v || true
    docker-compose --project-name go_nautobot -f ../docker-compose.yml down -v || true
}

trap cleanup EXIT

echo "Getting code generator project"

git clone https://github.com/swagger-api/swagger-codegen || true
cd swagger-codegen
git checkout 3.0.0
./run-in-docker.sh mvn package
cd ..

echo "Starting Nautobot instance"

docker-compose --project-name go_nautobot -f docker-compose.yml --env-file local_dev.env up --build --abort-on-container-exit --exit-code-from oapi

echo "Generating Go Bindings"
cd swagger-codegen
grep -q '-e _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999' ./run-in-docker.sh || sed -i '/-e MAVEN_CONFIG=\/var\/maven\/.m2/a \        -e _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999 \\' ./run-in-docker.sh
./run-in-docker.sh generate -i swagger.yaml -l go -o /gen/out/go-nautobot -DpackageName=go-nautobot

cp -arT ./out/go-nautobot/ ../../
cd ..

echo "Done"
