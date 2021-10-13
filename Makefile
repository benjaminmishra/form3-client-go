STATICCHECK = $(GOPATH)/bin/staticcheck
export PSQL_USER=root
export PSQL_PASSWORD=password
export PSQL_HOST=0.0.0.0
export PSQL_PORT=5432
export API_HOST=http://localhost:8080


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc

deps:
	go mod download

test: test.unit test.integration

test.unit: lint
	go test ./f3client/... -run=^Test_Unit_ -cover -v

test.integration: deps lint api.start
	go test ./f3client/... -p 1 -run=^Test_Integration_ -v -cover
	docker compose -f docker-compose.test.yml down --volumes

lint: fmt | $(STATICCHECK)
	go vet ./f3client/...
	$(STATICCHECK) ./f3client/...

fmt : deps
	go fmt ./f3client/...

doc :
	$(GODOC)
	godoc -http=:6060

api.start:
	docker compose -f docker-compose.test.yml up -d

api.stop:
	docker compose -f docker-compose.test.yml down --volumes


docker.cleanup:
	docker compose down
	docker image rmi form3-client-go_accountapi_client


test.cover:
	go tool cover -html=./test_results/unitcover.out
	go tool cover -html=./test_results/itcover.out