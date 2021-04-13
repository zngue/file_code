#!/usr/bin/env bash
proto=$1
protoc --go_out=plugins=grpc:. $proto