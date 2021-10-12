STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc

hello:
	echo "Hello World"

test: lint
	go test -cover -v

test.unit: lint
	go test ./f3client/... -cover -v

test.integration: lint
	go test ./f3client/... -tag=integration -v
	

lint: fmt | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

fmt : 
	go fmt ./...

doc :
	$(GODOC)
	godoc -http=:6060

docker.start.components:
	docker compose -f docker-compose.yml up

docker.stop.components:
	docker compose -f docker-compose.yml down --volumes