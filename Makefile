STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc


test: lint test.unit test.integration

test.unit: lint
	go test ./f3client/... -run=^Test_Unit -cover -v

test.integration: lint 
	go test ./f3client/... -p 1 -run=^Test_Integration -v -cover
	

lint: fmt | $(STATICCHECK)
	go vet ./f3client/...
	$(STATICCHECK) ./f3client/...

fmt :
	go fmt ./f3client/...

doc :
	$(GODOC)
	godoc -http=:6060

docker.start.components:
	docker compose -f docker-compose.yml up

docker.stop.components:
	docker compose -f docker-compose.yml down --volumes
