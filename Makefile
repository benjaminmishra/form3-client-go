STATICCHECK = $(GOPATH)/bin/staticcheck


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc

$(EXPORT):
	export PSQL_USER=root
	export PSQL_PASSWORD=password
	export PSQL_HOST=0.0.0.0
	export PSQL_PORT=5432
	export API_HOST=http://localhost:8080

$(UNSET):
	unset PSQL_USER
	unset PSQL_PASSWORD
	unset PSQL_HOST
	unset PSQL_PORT
	unset API_HOST


test: lint test.unit test.integration

test.unit: lint
	go test ./f3client/... -run=^Test_Unit_ -cover -v

test.integration: lint api.start
	$(EXPORT)
	go test ./f3client/... -p 1 -run=^Test_Integration_ -v -cover
	$(UNSET)
	docker compose -f docker-compose.test.yml down --volumes

lint: fmt | $(STATICCHECK)
	go vet ./f3client/...
	$(STATICCHECK) ./f3client/...

fmt :
	go fmt ./f3client/...

doc :
	$(GODOC)
	godoc -http=:6060

api.start:
	docker compose -f docker-compose.test.yml up -d

api.stop:
	docker compose -f docker-compose.test.yml down --volumes
