#!/usr/bin/env bash
export GOOS=linux
export CGO_ENABLED=0

go mod download

go mod tidy

cd tools/healthchecker;go build  -o healthchecker-linux-amd64;echo built `pwd`;cd ../..
go build -o echoservice-linux-amd64
echo "build  echoservice-linux-amd64  success!!!"

mv echoservice-linux-amd64 docker/
mv tools/healthchecker/healthchecker-linux-amd64 docker/
echo "mv  files  success!!!"