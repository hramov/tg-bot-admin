FROM golang:1.19-alpine as builder

RUN apk update
RUN apk upgrade
RUN apk --no-cache add -U ca-certificates
RUN apk add --no-cache bash

WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src/app /usr/src/app/src/app
COPY ./src/config /usr/src/app/src/config
COPY ./src/modules /usr/src/app/src/modules
COPY ./src/main.go /usr/src/app/src/main.go

COPY ./data/ /usr/src/app/data/

RUN go build -o ./bin/server ./src/main.go

EXPOSE 4000
CMD "./bin/server"