#!/bin/bash

server="./gin-apiserver"
let item=0
item=`ps -ef | grep $server | grep -v grep | wc -l`

if [ $item -eq 1 ]; then
	echo "The gin-apiserver is running, shut it down..."
	pid=`ps -ef | grep $server | grep -v grep | awk '{print $2}'`
	kill -9 $pid
fi

echo "Start gin-apiserver now ..."
make build
./build/pkg/cmd/gin-apiserver/gin-apiserver  >> gin-apiserver.log 2>&1 &
