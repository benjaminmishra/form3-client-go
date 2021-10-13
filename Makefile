STATICCHECK = $(GOPATH)/bin/staticcheck

# varaibles required to run the test suit using make only
# docker compose up has its own env vvars defined in compose file
export PSQL_USER=root
export PSQL_PASSWORD=password
export PSQL_HOST=0.0.0.0
export PSQL_PORT=5432
export API_HOST=http://localhost:8080


$(STATICCHECK):
	go get honnef.co/go/tools/cmd/staticcheck

$(GODOC):
	go get -v  golang.org/x/tools/cmd/godoc

# used to load all dependecies
deps:
	go mod download

# run this when you need to run the whole test suit
# does everything unit and integration statements do individually
test: test.unit test.integration

# run when you need only unit tests
# also generates test coverage results in test_results folder
test.unit: lint
	go test ./f3client/... -run=^Test_Unit_ -v -coverprofile=./test_results/unitcover.out


# run only for integration tests
# has 5 sec sleep to allow for the containes and apis to start up
# Note : In case you are running this on windows change "sleep 5" to "timeout 5"
test.integration: deps lint | api.start
	sleep 5
	go test ./f3client/... -p 1 -run=^Test_Integration_ -v -coverprofile=./test_results/itcover.out
	make api.stop

lint: fmt | $(STATICCHECK)
	go vet ./f3client/...
	$(STATICCHECK) ./f3client/...

fmt : deps
	go fmt ./f3client/...

# use this to see the documetation for this pkg
doc :
	$(GODOC)
	godoc -http=:6060

api.start:
	docker-compose -f docker-compose.test.yml up -d

api.stop:
	docker-compose -f docker-compose.test.yml down --volumes


docker.cleanup:
	docker-compose down
	docker image rmi form3-client-go_accountapi_client

# use this to analyse the test coverage results in html format
test.cover:
	go tool cover -html=./test_results/unitcover.out
	go tool cover -html=./test_results/itcover.out