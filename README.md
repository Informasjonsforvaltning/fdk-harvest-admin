# fdk-harvest-admin
A service that provide functionality to register and list data sources to be harvested.


## To run locally:
* install Golang >=1.17

* Tests:
```
// Run all tests
go test ./test
// Run tests with coverage report
go test -v -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./test
```
* Start the service:
```
// Start with go run
docker-compose up -d mongodb
go get
go run main.go
// Start with docker-compose
docker-compose up -d --build
```
* Check that service is running:
```
curl -X 'GET' \
  'localhost:8000/ping' \
  -H 'accept: application/json'
```