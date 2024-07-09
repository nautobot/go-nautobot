FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add --no-cache bash wget gcc musl-dev gcompat npm openjdk11-jre

RUN npm install @openapitools/openapi-generator-cli -g





COPY development/generate.sh ./
COPY development/oapi-config.yaml ./

CMD [ "./generate.sh" ]
