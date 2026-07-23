FROM golang:1.23-alpine

WORKDIR /app

ARG OPENAPI_GENERATOR_VERSION

RUN apk update && apk add --no-cache bash wget gcc musl-dev gcompat npm openjdk11-jre python3 py3-yaml

RUN test -n "${OPENAPI_GENERATOR_VERSION}" \
    && npm install @openapitools/openapi-generator-cli -g \
    && openapi-generator-cli version-manager set "${OPENAPI_GENERATOR_VERSION}"

CMD [ "/client/development/create-bindings/.create-bindings.sh" ]
