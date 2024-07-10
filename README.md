# Key Value Database API

Golang API using the [gin](https://gin-gonic.com/) framework.

Creates an empty in memory key-value database, using a Mutex to ensure that only one goroutine can access the database at a time to prevent data races and inconsistencieswhen multiple threads or processes are trying to read or write the same key simultaneously.

Explicitly sets the `GOMAXPROCS` environment variable to ensure that the maximum number of threads are available for the parallel execution of requests.  This is for the purpose of clarity as by default `GOMAXPROCS` is equal to the number of CPUs.

## Running the API
Start the server with: 

`go run cmd/main.go`

The API runs on port 8080.

Note that the database is empty and will need to be populated using with a PUT http request (see endpoints below).

## Endpoints

1. `/` GET - Returns a JSON list of all keys in the database.
2. `/<key>` GET - Returns the value for a given key, or a 404 if the key does not exist.
3. `/<key>` PUT - Updates the value for a given key.  Returns the updates key value pair or a 404 if the key does not exist.
4. `/<key>` DELETE - Deletes a key value pair for a given key.  Returns the deleted key, or a 404 if the key does not exist.

### Example curl requests to populate and query the database

## Testing

Tested using the `testing` package, using `t.Parallel()` to ensure that, where relevant, tests are run in parallel (by default tests run sequentially).  

Run the tests with:

`go test ./...`



