#!/bin/bash

set -e

gf build main.go -a amd64 -s linux -p ./temp
gf docker main.go -p -t xyhelper/auditlimit:latest
now=$(date +"%Y%m%d%H%M%S")
# 以当前时间为版本号
docker tag xyhelper/auditlimit:latest xyhelper/auditlimit:$now
docker push xyhelper/auditlimit:$now
echo "release success" $now
