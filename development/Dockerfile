FROM golang:1.18-alpine

WORKDIR /app

RUN apk update \
    && apk add gcc musl-dev gcompat

# TODO: try v2?
RUN go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
# RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

COPY development/generate.sh ./
COPY development/oapi-config.yml ./

CMD [ "./generate.sh" ]
