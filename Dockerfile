FROM golang:1.12.2-alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD main.go ./

RUN go build -o main

CMD ./main


