FROM golang:1.18-alpine

WORKDIR /app

RUN apk update \
    && apk add gcc musl-dev gcompat

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

COPY development/generate.sh ./

CMD [ "./generate.sh" ]
