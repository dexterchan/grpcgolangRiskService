#!bin/bash
eval export $(go env | grep GOPATH)
PATH=$GOPATH/bin:$PATH
mkdir -p pkg/api/v1
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 todo-service.proto
mkdir -p pkg/api/risk
protoc --proto_path=api/proto/risk --proto_path=third_party --go_out=plugins=grpc:pkg/api/risk risk.proto
