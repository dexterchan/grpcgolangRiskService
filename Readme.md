https://github.com/golang-standards/project-layout?source=post_page-----daebb36a97e9----------------------

https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9

go test ./pkg/api/...

cd cmd/server
go build .

## Run server without log for first implementation lab 1

```
server -grpc-port=9090 -db-host=localhost:3306 -db-user=agileintelligence -db-password=password -db-schema=ppmtcourse
```

## Run server with log

```
server -grpc-port=9090 -http-port=8080 -db-host=localhost:3306 -db-user=agileintelligence -db-password=password -db-schema=ppmtcourse
```

## Run server with trace

```
server -grpc-port=9090 -http-port=8080 -db-host=localhost:3306 -db-user=agileintelligence -db-password=password -db-schema=ppmtcourse -log-level=-1 -log-time-format=2006-01-02T15:04:05.999999999Z07:00
```

### Run grpc client

```
cd cmd/client-grpc
client-grpc -server=localhost:9090
```

### Run restful client

```
client-rest -server=http://localhost:8080
```
