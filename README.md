# quotes-service-go
A simple Go CRUD API for Quotes written in Go. For the sake of simplicity this is only using an in memory slice to hold the quote values. Quotes are immutable in this application.

## Quick Start 

```
# install golang
brew install golang

# install the golangci linter
# more details: https://golangci-lint.run/
brew install golangci-lint
```

### Downloading dependencies
```
go mod download
```

### Building Locally
```
# Make build
go build main.go

# Run application
go run main.go
```

### Running unit tests
```
# Runs all the tests present in the tests directory
go test -v ./tests/*.go
```

## API Documentation
The API supports the following requests:
- Health Check
    - Method: GET
    - URL: /health
    - Responses: 
        - 200: Application is healthy. 

-   Create quote
    -   Method: POST
    -   URL: /api/v1/quotes
    -   Parameters:
        -   id: The id you wish to insert (ex: "3") (required)
        -   quote: Quote in quection (ex: 'Live life')(required)
        -   author_name: Person who the quote belongs to (ex: 'Morgan Freeman') (required)
    -   Responses:
        -   201: Quote created successfully


-   Get a list of all the quotes
    -   Method: GET
    -   URL: /api/v1/quotes
    -   Responses:
        -   200: Get all quotes successfully


-   Read details of a quote
    -   Method: GET
    -   URL: /api/v1/quotes/{id}
    -   {id}: The id of the quote in question
    -   Responses:
        -   200: Got quote successfully
        -   404: Quote does not exist


-   Delete a quote
    -   Method: DELETE
    -   URL: /api/v1/quotes/{id}
    -   {id}: The id of the quote in question
    -   Responses:
        -   200: Quote deleted successfully
        -   404: Quote does not exist