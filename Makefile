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

test.unit: lint
	go test -tags=unit_tests -v

test.integration: lint
	go test -tags=integration_tests -v