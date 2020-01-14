#!/bin/bash
SRC_DIR="./"
GO_DIR="./proto/"
PY_DIR="./python/"
python3 -m grpc_tools.protoc --proto_path=$SRC_DIR  --python_out=$PY_DIR --grpc_python_out=$PY_DIR $1
