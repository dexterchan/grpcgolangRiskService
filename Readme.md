https://github.com/golang-standards/project-layout?source=post_page-----daebb36a97e9----------------------

https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9

go test ./pkg/api/...


cd cmd/server
go build .
server -grpc-port=9090 -db-host=localhost:3306 -db-user=agileintelligence -db-password=password -db-schema=ppmtcourse

cd cmd/client-grpc
client-grpc -server=localhost:9090