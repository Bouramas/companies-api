# Companies API

This is a sample REST API with basic CRUD operations available.

Implementation details:
* GoLang 1.19
* MySQL
* Docker

## Steps to run locally

### DB
Run `docker-compose up -d --build mysql`

### API
1. Set the MYSQL_DSN env var used by the API to connect to the DB:
   `export MYSQL_DSN="root:password@tcp(mysql:3306)/the_company_db?parseTime=true&sql_mode=NO_ZERO_DATE"`
2. Modify `GOOS` and `GOARCH` based on your machine in the `Makefile`
3. Run `make build`
4. Run `./companies-api`

You can also use the Postman collection in the [docs](docs) folder
.


### TODO - FIXME:

// Run with Dockerfile and Docker-compose
// Add API documentation
// Add Unit Tests

// Docker run
// Not working:

// Build the image
docker build -t c-api-image .

// Run the image
docker run -p 8080:8080 c-api-image


