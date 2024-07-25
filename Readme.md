# Go Server Template

A simple example of how to start an API in Golang using the [Gin framework](https://gin-gonic.com/#td-block-1).

## Running

To build the project, run the command:

```
make setup
```

To run the api use the command:

```
go run main.go
```

The api will be available on the port `8380`

## Request Example

```
curl --location 'http://localhost:8380/task' \
--header 'Content-Type: application/json' \
--data '{
    "text": "Hello, World!"
}'
```
