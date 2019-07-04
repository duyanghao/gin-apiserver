#!/bin/bash

server="./GinApiServer"
let item=0
item=`ps -ef | grep $server | grep -v grep | wc -l`

if [ $item -eq 1 ]; then
	echo "The GinApiServer is running, shut it down..."
	pid=`ps -ef | grep $server | grep -v grep | awk '{print $2}'`
	kill -9 $pid
fi

echo "Start GinApiServer now ..."
make build
./build/pkg/cmd/GinApiServer/GinApiServer  >> GinApiServer.log 2>&1 &
