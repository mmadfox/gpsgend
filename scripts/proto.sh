#!/bin/bash

rootPath=$(pwd)

docker run --rm -v $rootPath:$rootPath -w $rootPath github.com/mmadfox/gpsgend/protoc:latest \
 --proto_path=.                                                                              \
 --go_out=./gen --go_opt=paths=source_relative                                               \
 --go-grpc_out=./gen --go-grpc_opt=paths=source_relative                                     \
  proto/gpsgend/v1/types.proto                                                               \
  proto/gpsgend/v1/generator_service.proto                                                   \
  proto/gpsgend/v1/tracker_service.proto