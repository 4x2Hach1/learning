### build ####################################################
FROM golang as build

ENV GOTOOLCHAIN=auto
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
WORKDIR /app

COPY . .
RUN go build -o ./api

#### deploy ###################################################
FROM gcr.io/distroless/static-debian11 as deploy

COPY --from=build /app/api /api

EXPOSE 8080
ENTRYPOINT [ "/api" ]
