FROM golang:1.18-alpine

WORKDIR /app

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

COPY development/generate.sh ./

CMD [ "./generate.sh" ]
