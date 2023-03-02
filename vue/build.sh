#!/bin/bash
cd `dirname $0`

image=andreleesss/tods_server
tag=latest
echo $image:$tag

sudo docker stop tods_web
sudo docker rm tods_web
sudo docker rmi $image:$tag
rm -rf ./dist/*

set -e
npm run build
sudo docker build -t $image:$tag .
rm -rf ./dist/*
