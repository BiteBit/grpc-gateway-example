#! /bin/bash

go build -o protoc-gen-box cmd/protoc-gen-box/*
mv protoc-gen-box ~/go/bin/protoc-gen-box

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --box_out=logtostderr=true:. \
  api/api.proto
