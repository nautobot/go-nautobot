FROM alpine:latest

WORKDIR /app

RUN apk update && apk add --no-cache bash wget

CMD [ "/client/development/get-api/.get-api.sh" ]
