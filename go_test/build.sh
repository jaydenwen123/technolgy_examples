#!/bin/bash
if [ $# -eq 2 ];then
	GOOS=$1
	GOARCH=$2
fi
echo "[INFO] GOOS:"$GOOS",GOARCH:"$GOARCH
go test -c 
