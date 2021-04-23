#!/usr/bin/env bash

protoc --go_out=plugins=grpc:../pbmodel *.proto

#生成所有的pb文件