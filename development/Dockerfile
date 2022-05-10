FROM golang:1.18-alpine

WORKDIR /app

RUN apk update \
    && apk add gcc musl-dev gcompat

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0

COPY development/generate.sh ./
COPY development/oapi-config.yml ./

CMD [ "./generate.sh" ]
