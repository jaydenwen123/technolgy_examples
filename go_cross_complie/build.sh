#!/bin/bash
obj="go_cross_complie"
if [ $# -eq 2 ];then
	GOOS=$1
	GOARCH=$2
	echo "[INFO] GOOS:"$GOOS",GOARCH:"$GOARCH
#	交叉编译，需要按照下面写
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o ./${GOOS}/${obj}
else
  go build -o ${obj}
fi

