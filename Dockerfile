FROM golang:alpine as builder

RUN apk add build-base
ENV GO111MODULE=on
ADD . /go/src/github.com/ginuerzh/gost/
WORKDIR /go/src/github.com/ginuerzh/gost/cmd/gost
RUN go install -buildmode=pie -trimpath

FROM alpine:latest

COPY --from=builder /go/bin/gost /usr/local/bin/gost
ENTRYPOINT ["/usr/local/bin/gost"]

