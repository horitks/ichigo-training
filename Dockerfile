FROM golang:1.18.1-alpine3.15 as builder

WORKDIR /go/src
COPY src .

RUN go build main.go

