FROM golang:1.11-alpine as builder

WORKDIR /go/src/upspinserver
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

RUN go get github.com/palager/gcp/cmd/upspinserver-gcp

FROM alpine:latest
COPY --from=builder /go/bin/upspinserver-gcp /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/bin/upspinserver-gcp"]

