FROM golang:1.22 as build

RUN apt-get update && apt-get install unzip

WORKDIR /app

ARG PROTOC_VERSION=3.20.3
RUN curl -Lo protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip" && \
  unzip -q protoc.zip bin/protoc 'include/*' -d /usr/local

RUN go install github.com/go-task/task/v3/cmd/task@latest
