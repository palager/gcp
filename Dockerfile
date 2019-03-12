FROM golang:1.11-alpine as builder

WORKDIR /go/src/upspinserver
RUN apk update \
        && apk upgrade \
        && apk add --no-cache ca-certificates git \
        && update-ca-certificates 2>/dev/null || true


RUN go get github.com/palager/gcp/cmd/upspinserver-gcp

FROM alpine:latest
COPY --from=builder /go/bin/upspinserver-gcp /bin

CMD ["/bin/upspinserver-gcp"]

