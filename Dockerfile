FROM golang:1.22-alpine AS build-env
 
ENV APP_NAME=fdk-harvest-admin
ENV CMD_PATH=main.go
 
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

FROM alpine:latest

ENV APP_NAME=fdk-harvest-admin
ENV GIN_MODE=release

COPY --from=build-env /$APP_NAME .

EXPOSE 8080

CMD ["sh", "-c", "./$APP_NAME"]
