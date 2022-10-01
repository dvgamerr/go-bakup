FROM golang:alpine as builds

RUN go install github.com/dvgamerr/go-bakup

ENTRYPOINT [ "go-bakup" ]