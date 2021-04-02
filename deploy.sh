#!/usr/bin/env bash

case "$1" in
    build)
        if [ "$2" != "linux" ] && [ "$2" != "darwin" ] && [ "$2" != "windows" ]; then
          echo "不支持的编译环境 ($2)"
          exit 1
        fi
        echo "build ($2)"
        export CGO_ENABLED=0
        export GOOS=$2 # linux darwin windows
        export GOARCH=amd64

        go build
    ;;
    *)
        echo "deploy ($1)"
    ;;
esac
