ARG VERSION
FROM golang:$VERSION
RUN mkdir -p /go/src/github.com/absaleb/notifier
WORKDIR /go/src/github.com/absaleb/notifier
COPY . .
RUN export GO111MODULE=on CGO_ENABLED