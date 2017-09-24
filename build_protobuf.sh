#!/bin/sh

protoc --go_out=plugins=grpc:. core/*.proto
protoc --go_out=plugins=grpc:. stars/*.proto
