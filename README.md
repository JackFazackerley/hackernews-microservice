# Hackernews Microservice

Simple project to demo microservice architecture in Go by pulling the top stories from hackernews and storing them in a
db.

This isn't tested, it should be...

## What
This project has two parts; [Core](core) and [GRPC](grpc). Core demonstrates how to get data from Hackernews, store 
said data, and retrieve via a custom API. The GRPC service demonstrates the same functionality but uses GRPC to retrieve 
the data instead.

Both services have a REST API. The Core REST API directly communicates with the database, however within the GRPC 
service, the REST API communicates with a GRPC server, which in turn communicates with the database.

## API
Each service has the same endpoints to retrieve data:

* /All - returns all items regardless of type
* /Jobs - returns only type `job`
* /Stories - returns only type `story`

## How to run
### Core
```shell
$(cd core && docker-compose up) 
```

### GRPC
```shell
$(cd grpc && docker-compose up)
```

## How to destroy
### Core
```shell
$(cd core && docker-compose down) 
```

### GRPC
```shell
$(cd grpc && docker-compose down)
```

## Note
You'll notice that there is a fair amount of code duplication in this project between core and grpc. This is 
intentional, purely for showing the difference between them.