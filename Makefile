all: test build
build:
	go build -o upspinserver-gcp -v ./cmd/upspinserver-gcp
	go build -o keyserver-gcp -v ./cmd/keyserver-gcp
test:
	go test -v ./...
clean:
	go clean ./... && rm ./upspinserver-gcp
