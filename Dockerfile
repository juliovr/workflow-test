FROM golang:alpine

ARG app_env
ENV APP_ENV $app_env

ADD . /app
WORKDIR /app

RUN apk add git
RUN go get -d -v ./...
RUN go build main.go

CMD ["/app/main"]

EXPOSE 8003
