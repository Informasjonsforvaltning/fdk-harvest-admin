# FDK HARVEST ADMIN

This application provides an API to register and list data sources to be harvested.

For a broader understanding of the system’s context, refer to
the [architecture documentation](https://github.com/Informasjonsforvaltning/architecture-documentation) wiki. For more
specific context on this application, see the **Harvesting** subsystem section.

## Getting Started

These instructions will give you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

Ensure you have the following installed:

- Go
- Docker

Clone the repository.

```sh
git clone https://github.com/Informasjonsforvaltning/fdk-harvest-admin.git
cd fdk-harvest-admin
```

#### Start MongoDB and RabbitMQ

```sh
docker-compose up -d
```

If you have problems starting kafka, check if all health checks are ok. Make sure number at the end (after 'grep')
matches desired topics.

#### Install required dependencies

```sh
go get
```

#### Start application

```sh
go run main.go
```

### API Documentation (OpenAPI)

The API documentation is available at ```spec/fdk-harvest-admin.yaml```.

### Running tests

```sh
go test ./test
```

To generate a test coverage report, use the following command:

```sh
go test -v -race -coverpkg=./... -coverprofile=coverage.txt -covermode=atomic ./test
```
