FROM golang:1.15 AS builder2
ENV GOPROXY "https://goproxy.cn"
ENV GO111MODULE on
USER root
WORKDIR /root

COPY . /cloud-element
COPY go.mod go.sum ./
RUN go mod download
RUN apt-get update && apt-get install -y libhwloc-dev && go build