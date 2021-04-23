#!/usr/bin/env bash
proto=$1
protoc --go_out=plugins=grpc:. $proto


##生成单个pb文件