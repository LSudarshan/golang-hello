FROM golang:1.7.3

RUN mkdir -p /go/src/github.com/lsudarshan/hello
WORKDIR /go/src/github.com/lsudarshan/hello

CMD ["/go/bin/hello"]

COPY . /go/src/github.com/lsudarshan/hello/
RUN go install github.com/lsudarshan/hello/utils
RUN go install github.com/lsudarshan/hello/