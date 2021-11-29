FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-ueslessfacts

FROM alpine:3.15.0
WORKDIR /app
COPY --from=build-env /src/gowon-ueslessfacts /app/
ENTRYPOINT ["./gowon-ueslessfacts"]
