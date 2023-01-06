FROM golang:1.19-alpine as builder

RUN apk update
RUN apk upgrade
RUN apk --no-cache add -U ca-certificates
RUN apk add --no-cache bash

WORKDIR /usr/src/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src/ /usr/src/app/src/
COPY ./data/ /usr/src/app/data/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o /usr/bin/server ./src/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app/bin/server /usr/bin/server
EXPOSE 4000
ENTRYPOINT ["/usr/bin/server"]

