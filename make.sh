#!/bin/bash

set -ex

CUR_DIR=$PWD
SRC_DIR=$PWD
BIN_DIR=$PWD/bin
SERVICE_DIR=$PWD/services/
FLAGS=-race
export GOBIN=$BIN_DIR

go generate ./...

cd $SRC_DIR
go vet ./...

cd $SERVICE_DIR
for plugin_name in `ls -l | grep '^d' | awk '{print $9}' | grep -v 'internal'`; do
    go build $FLAGS -buildmode=plugin -o $BIN_DIR/$plugin_name.so ./$plugin_name;
done
cd $SRC_DIR
go install $FLAGS .

case $1 in
    "docker") docker build -t go-xserver . ;;
    "start")
        cd $BIN_DIR
        ;;
    ?);;
esac

cd $CUR_DIR
