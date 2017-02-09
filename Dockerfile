FROM golang:1.7.5-alpine

WORKDIR /go/src/github.com/weaviate/weaviate

RUN apk add --no-cache git bash

COPY weaviate.go /go/src/github.com/weaviate/weaviate/
COPY core /go/src/github.com/weaviate/weaviate/core/

RUN ls -al

RUN go get -d -v
RUN go install

CMD /go/bin/weaviate

EXPOSE 80
