# f3client
This library wraps the [form3 v1 apis](https://api-docs.form3.tech/api.html) into a simple reusable client library. Right now this library only supports the Account (Create, Fetch, Delete) api.  

Currently the f3client libray requires go version 1.17.2 or greater.

The structure of the library is inspired by the followig projects and borrows some ideas from them.
- [go-github](https://github.com/google/go-github)
- [twillio-go](https://github.com/kevinburke/twilio-go)
- [godo](https://github.com/digitalocean/godo)

Also it follows advice from the following blog posts about semantics
- [Dave Cheny's blog about functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Josh Michielsen's article about api clients for humans](https://blog.gopheracademy.com/advent-2019/api-clients-humans/)
- [George Shaw's two part article on integration testing with docker](https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html)

>Note for reviewer: I am fairly new to using go and wirting open source libraries in general. Hence have heavily relied on blogs metioned above.

## Installation

f3client is compatable with module mode in modern Go release. In order to get this module add the follwoing to your code and run go get.

``` go
import github.com/benjaminmishra/form3-client-go/v0.0.1
```
Alternatively you can just run the following in your project root directory and start importing in your code.
```bash
go get github.com/benjaminmishra/form3-client-go/v0.0.1
```

## Usage
Import the module 
``` go
import github.com/benjaminmishra/form3-client-go/v0.0.1/f3client // with go modules enabled (GO111MODULE=on or outside GOPATH
import github.com/benjaminmishra/form3-client-go/f3client // with go modules disabled
```

Construct a new f3clinet and then use the various services on the client to access diffrent parts of the api. For example :
``` go
// create new f3client, with default options
c, err := f3client.NewClient()

// call the fetch function on the client.Account service
// here the accountId is mandatory
//
// this returns a account object and nil error
// in case of error account is nil and err is non-nil
account, err := client.Accounts.Fetch(ctx, accountId)
```

The create api needs an pointer to an instance of f3clinet.account struct to be passed in order for get the full details of the acocunt. Note that in all cases the account id and/or organisation id is mandatort.
For example :

```go
account := f3client.Account{
	ID:             accountId,
	OrganisationID: orgId,
	Attributes: f3client.AccountAttributes{
		Country: "GB",
		Name:    []string{"jane doe", "john doe"},
    },
}

err = c.Accounts.Create(bctx, &account)
```
More details on each fields can be found in the form3 api documentation for individual apis.

## Tests
I have relied heavily on makefile to automate the running of both integration and unit tests for this module. You can run the tests both directly from your system or using docker compose up command. The steps for each of them is described below.

Note that you need to have GNU make installed on your system. Also docker and docker compose needs to 

### Running tests directly on your system:
*__Prerequisites__ : GNU Make, go version 1.17 and higher, docker , docker compose, docker engine running*
1. Clone this repo on your system.
2. Open terminal or command line and cd into the root directory of the repo and run.
   ``` bash
   make test
   ```
3. This will trigger both the unit and integration tests. 
   - Unit tests would run and produce a coverage report in test_results folder also in the command line
   - Integration tests would launch the only supporting docker containers using docker compose up, run the tests and shut donw the containers. This also produces coverage results in the same folder.
4. You can also run the unit tests and integration tests seperately for that you need to type in the following commands
   ```bash
   make test.unit  # for running only unit tests
   make test.integration # for running only integration, also sets the neccery env variables 
   ```

>Note : When running tests the code is first linted and staticchcker is run on it first. In case you have >made modifications to the code that violate the go rules of formatting then the tests will fail to run.

>Also this codebase uses the docker-compose.yml file provided and builds up on it. It has two >docker-compose files i.e. docker-compose.yml and docker-compose.test.yml. To test using docker compose up >it uses the docker-compose.yml file. But to test direcly on system it uses docker-compose.test.yml.

### Running tests using docker compose up
*__Prerequisites__ : GNU Make, go version 1.17 and higher, docker , docker compose, docker engine running*
1. Clone the repo on your system and cd into the root directory.
2. Run ```docker compose up```
3. This will do the following  in the background
   - Launch all supporting docker containers and launch a container that wraps the client codebase(see Dockerfile included in the code)
   - Run the test code the in the client , both unit tests and integration tests against the containerized form3 account apis
4. Once tests are run the container that packages the client code exists 
5. This also produces coverage results in the test_results folder in .out format.
6. To cleanup you can either 
   - run ```docker compose down```
   - run ```make docker.cleanup``` in the root directory of the codebase.

## Versioning
This library follows semver at a high level. At this moment this is labelled as minor version 1.


## Features yet to be implemened and issues

Obviously this codebase lacks some features and has some issues that would be good to have in case this is deployed in a full production scenarios. Some of them are listed below.

### Features

- *Rate limiting* : This library does not implement any rate limiting. 
- *Better Error Handling* : The error handling side of the code needs to be improved so as to provide the user with meaningful errors.
- *Authentication Support* : No support for any kind of authetication. 
- *Context support* - At this moment there is no context support , although the functions do require context to be passed in, but its not being handled anywhere.
- *Better testing* - You can never test enough. Right now this repo only has 88% unti test coverage and 77% integration test coverage.


### Issues
- *Error Handling* - There are certain senarios where the clinet fails to handle the errors generated by the apis correctly.


## License :
This codebase follows MIT License. More about that in [LICENSE](/LICENSE) file.
