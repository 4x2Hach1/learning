### develop ####################################################
FROM golang:1.21.4 as build

RUN apt-get update && apt-get install unzip

WORKDIR /app

ARG PROTOC_VERSION=3.20.3
RUN curl -Lo protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip" && \
  unzip -q protoc.zip bin/protoc 'include/*' -d /usr/local

### build ####################################################
# FROM golang:1.20.5 as build

# ENV CGO_ENABLED=0
# WORKDIR /app

# COPY . .
# RUN go mod download
# RUN go build -o ./api

#### deploy ###################################################
# FROM gcr.io/distroless/static-debian11 as deploy

# COPY --from=build /app/api /api

# EXPOSE 8080
# ENTRYPOINT [ "/api" ]
