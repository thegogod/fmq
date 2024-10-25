fmt:
	gofmt -w ./

test:
	go clean -testcache
	go test ./... -cover

run: server.run

server.run:
	FMQ_PLUGINS=mqtt $(MAKE) -C server run
