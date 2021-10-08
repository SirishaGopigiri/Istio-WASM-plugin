FROM golang:alpine AS build-env
RUN mkdir /go/src/app && apk update && apk add git
ADD app.go /go/src/app/
ADD go.mod /go/src/app/
ADD go.sum /go/src/app/
WORKDIR /go/src/app
RUN go build -o app .
WORKDIR /go/src/app
ENTRYPOINT [ "./app" ]
