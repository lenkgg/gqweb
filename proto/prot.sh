#!/bin/bash
SRC_DIR=./
GO_DIR="./proto/"
protoc -I=$SRC_DIR --go_out=plugins=grpc:$GO_DIR $1
