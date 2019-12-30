#!/bin/bash
export GOOS=linux
export GOARCH=amd64
go build -o blog ./main.go &
wait
scp -r blog root@hk:/root/blogs