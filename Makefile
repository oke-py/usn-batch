.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/usn-batch usn/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --stage dev --verbose

deployprod: clean build
	sls deploy --stage prod --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
