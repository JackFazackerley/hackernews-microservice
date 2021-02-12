# Hackernews Microservice

Simple project to demo microservice architecture in Go by pulling the top stories from hackernews and storing them in a
db.

This isn't tested, I didn't want to, simply a mess around.

## How to run
```
docker-compose build && docker-compose up
go run cmd/presenter/main.go
```

To hit different endpoints within the presenter changing the path within `cmd/presenter/main.go` is required.

## How to destroy

```
docker-compose down
```