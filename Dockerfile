FROM golang:1.11-alpine as builder

WORKDIR /go/src/upspinserver
RUN apk --no-cache add git
RUN go get github.com/palager/gcp/cmd/upspinserver-gcp

FROM alpine:latest
COPY --from=builder /go/bin/upspinserver-gcp /bin

CMD ["/bin/upspinserver-gcp"]

