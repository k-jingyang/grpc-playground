## Introduction

Confusion was created by the existence of the two protobuf repositories:

1. [https://github.com/golang/protobuf](https://github.com/golang/protobuf), containing the first major version
1. [https://github.com/protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go) containing the second major version

To help beginners such as myself, this README serves as a guide to create a

- Golang
- Hello world project (Based on [grpc's own](https://github.com/grpc/grpc-go/tree/master/examples/helloworld))
- Using gRPC
- With the first major version of protobuf (The second major version does not support the creation of gRPC code)

**Last Updated: May 2020**

## Installation of protocol buffer compiler

Follow installation guide [here](https://grpc.io/docs/protoc-installation/)

## Installation of protoc-gen-go plugin for protoc

`$ go get github.com/golang/protobuf/protoc-gen-go`

## Installation of gRPC for Go

`$ go get google.golang.org/grpc`

## To compile .proto files for gRPC on Go

`$ protoc --go_out=paths=source_relative,plugins=grpc:. playground/playground.proto`
