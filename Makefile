.PHONY: build clean deploy

all: build fix vet fmt test lint tidy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/usn-batch usn/main.go

lint:
	golangci-lint run ./...

test:
	go test -v -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --stage dev --verbose

deployprod: clean build
	sls deploy --stage prod --verbose

fix:
	go fix ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

vet:
	go vet ./...
