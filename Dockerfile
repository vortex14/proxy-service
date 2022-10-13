FROM golang:1.17-alpine as build-stage

RUN mkdir -p /app

WORKDIR /app

COPY . /app
RUN go mod download
RUN  go mod vendor

RUN go build -o pserver main.go

FROM alpine:latest

WORKDIR /

COPY --from=build-stage /app .