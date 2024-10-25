fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

run:
	LOG_PREFIX=fmq/server go run ./...

build:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
