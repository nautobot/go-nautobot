FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add --no-cache bash wget gcc musl-dev gcompat openjdk11-jre

COPY generate.sh .

CMD [ "./generate.sh" ]
