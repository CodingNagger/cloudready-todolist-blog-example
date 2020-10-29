# cloudready-todolist-blog-example

Just a repository used as an example for a blog post I'm writing. Also using it to try and find my Go project structure style. 
Lots of inspiration taken from [https://github.com/katzien/go-structure-examples](https://github.com/katzien/go-structure-examples).

Basically a todolist server to create, list and complete tasks.

## How to run?

### Running tests

```bash
go test ./...
```

### Run the server

```bash
# Startup the local stack
TMPDIR=./deployment/local/tmp PORT_WEB_UI=8181 docker-compose -f ./deployment/localstack/docker-compose.yml up

# Create table
AWS_ACCESS_KEY_ID=example AWS_SECRET_KEY=example go run -v ./cmd/init/table -region us-east-1 -endpoint http://localhost:4566

# Build application
go build -v ./cmd/server/

# Start application
AWS_ACCESS_KEY_ID=example AWS_SECRET_KEY=example ./server -region us-east-1 -endpoint http://localhost:4566
```
