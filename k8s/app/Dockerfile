### build ####################################################
FROM golang:1.23.2 as build

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o ./api

#### deploy ###################################################
FROM gcr.io/distroless/static-debian11 as deploy

COPY --from=build /app/api /api

EXPOSE 8080
ENTRYPOINT [ "/api" ]
