#!/bin/bash

rootPath=$(pwd)

docker run --rm -v $rootPath:$rootPath -w $rootPath github.com/mmadfox/gpsgend/protoc:latest \
 --proto_path=.                                      \
 --go_out=. --go_opt=paths=source_relative           \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  internal/transport/grpc/types.proto internal/transport/grpc/service.proto