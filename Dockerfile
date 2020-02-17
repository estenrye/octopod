FROM golang:1.13.5-alpine3.11 AS build

COPY . /go/src
WORKDIR /go/src/cmd/server

RUN GOOS=linux   GOARCH=amd64 go build -o ./octopod

FROM alpine:3.11.3
COPY --from=build /go/src/cmd/server/octopod /octopod
ENTRYPOINT [ "/octopod" ]