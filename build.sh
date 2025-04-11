#!/bin/bash
# 构建 Linux/amd64
GOOS=linux GOARCH=amd64 go build -o bin/myapp-linux-amd64