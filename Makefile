all: test build
build:
	go build -o upspinserver-gcp -v ./cmd/upspinserver-gcp
test:
	go test -v ./...
clean:
	go clean ./... && rm ./upspinserver-gcp
