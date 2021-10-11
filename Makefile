STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc

hello:
	echo "Hello World"

test.unit: lint
	go test ./f3client/... -cover -v

test.integration: lint
	go test ./e2etest/... -cover -v

lint: fmt | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

fmt : 
	go fmt ./...


doc :
	$(GODOC)
	godoc -http=:6060