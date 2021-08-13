FROM golang:1.15 AS builder2
ENV GOPROXY "https://goproxy.cn"
ENV GO111MODULE on
USER root
WORKDIR /root

ADD ../cloud-element .
COPY go.mod go.sum ./
RUN go mod download