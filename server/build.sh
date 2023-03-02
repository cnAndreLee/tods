#!/bin/bash
cd `dirname $0`

image=andreleesss/tods_server
tag=latest
echo $image

sudo docker stop tods_server
sudo docker rm tods_server
sudo docker rmi $image:$tag
rm tods_server

set -e
go build
sudo docker build -t $image:$tag .
rm tods_server
