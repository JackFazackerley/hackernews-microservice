ARG go_version=1.15

#User
FROM alpine:3.13.1 as user
ARG uid=10001
ARG gid=10001
RUN echo "scratchuser:x:${uid}:${gid}::/home/scratchuser:/bin/sh" > /scratchpasswd

#Certs
FROM alpine:3.13.1 as certs
RUN apk add -U --no-cache ca-certificates

# BUILDER
FROM golang:${go_version}-alpine as build

WORKDIR /code/
ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/

RUN GOOS=linux CGO_ENABLED=0 GOGC=off GOARCH=amd64 go build -o ./bin/consumer ./cmd/consumer
RUN GOOS=linux CGO_ENABLED=0 GOGC=off GOARCH=amd64 go build -o ./bin/api ./cmd/api
RUN GOOS=linux CGO_ENABLED=0 GOGC=off GOARCH=amd64 go build -o ./bin/grpc ./cmd/grpc

FROM scratch as consumer
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=user /scratchpasswd /etc/passwd
COPY --from=build /code/bin/consumer .
ENTRYPOINT ["./consumer"]

FROM scratch as api
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=user /scratchpasswd /etc/passwd
COPY --from=build /code/bin/api .
ENTRYPOINT ["./api"]

FROM scratch as grpc
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=user /scratchpasswd /etc/passwd
COPY --from=build /code/bin/grpc .
ENTRYPOINT ["./grpc"]