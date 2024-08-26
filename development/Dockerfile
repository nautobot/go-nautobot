FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add --no-cache bash wget gcc musl-dev gcompat npm openjdk11-jre python3 py3-yaml

RUN npm install @openapitools/openapi-generator-cli -g

CMD [ "/client/development/generate.sh" ]
