all: test build
build:
	go build -o upspinserver-gcp -v ./cmd/upspinserver-gcp
	go build -o keyserver-gcp -v ./cmd/keyserver-gcp
test:
	go test -v ./...
clean:
	go clean ./... && rm ./upspinserver-gcp

push-upspinserver:
	docker build -t upspin-upspinserver-gcp -f Dockerfile.upspinserver-gcp .
	docker tag upspin-upspinserver-gcp gcr.io/cosmocr-at-upspin/upspin-upspinserver-gcp:latest
	docker push gcr.io/cosmocr-at-upspin/upspin-upspinserver-gcp:latest

push-keyserver:
	docker build -t upspin-keyserver-gcp -f Dockerfile.keyserver-gcp .
	docker tag upspin-keyserver-gcp gcr.io/cosmocr-at-upspin/upspin-keyserver-gcp:latest
	docker push gcr.io/cosmocr-at-upspin/upspin-keyserver-gcp:latest

push: push-upspinserver push-keyserver

