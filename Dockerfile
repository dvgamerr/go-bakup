FROM golang:alpine as builds

RUN go install github.com/dvgamerr/go-bakup@latest

ENTRYPOINT [ "go-bakup" ]