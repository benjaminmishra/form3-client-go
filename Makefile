STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck

hello:
	echo "Hello World"

test: lint
	go test -cover -v

lint: fmt | $(STATICCHECK)
	go vet ./...
	$(STATICCHECK) ./...

fmt : 
	go fmt ./...

tidy :
	go mod tidy

internal_test: lint
	go test -tags=internal_tests -v