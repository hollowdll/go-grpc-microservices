#!/bin/bash

# This script generates source codes from .proto files
# using Protobuf compiler protoc.

protoc --go_out=. --go-grpc_out=. \
  api/pb/**/*.proto
