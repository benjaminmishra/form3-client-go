STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

hello:
	echo "Hello World"

test: lint
	go test ./... -cover -v

lint: fmt | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

fmt : 
	go fmt ./...


doc : 
	godoc -http=:6060