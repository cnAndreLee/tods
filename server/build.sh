#!/bin/bash
cd `dirname $0`

sudo docker stop tods_server
sudo docker rm tods_server
sudo docker rmi tods_server
rm tods_server

set -e
go build
sudo docker build -t tods_server .
rm tods_server
