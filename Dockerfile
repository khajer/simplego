
#FROM golang:alpine
FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /app/simpleserver
COPY runserver . 

ENTRYPOINT /app/simpleserver/runserver
