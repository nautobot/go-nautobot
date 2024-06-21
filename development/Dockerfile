FROM golang:1.20-alpine

WORKDIR /app

RUN apk update \
    && apk add gcc musl-dev gcompat

RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

COPY development/generate.sh ./
COPY development/oapi-config.yml ./

CMD [ "./generate.sh" ]
