FROM golang:1.8.0-alpine

RUN apk add --no-cache git bash

COPY cmd /go/src/github.com/weaviate/weaviate/cmd
COPY models /go/src/github.com/weaviate/weaviate/models
COPY restapi /go/src/github.com/weaviate/weaviate/restapi

WORKDIR /go/src/github.com/weaviate/weaviate/cmd/weaviate-server

RUN go get -d -v
RUN go install

CMD ["/go/bin/weaviate-server", "--scheme", "http", "--port", "80"]

EXPOSE 80
