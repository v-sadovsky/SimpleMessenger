
###  Generate all necessary vendor options
make vendor

### Generate code according to API schema (also lint and format schema)
make generate

### Build client and server
make build

### Run Swagger UI service (will be available via http://localhost:8084/swaggerui endpoint)
make swagger-ui
