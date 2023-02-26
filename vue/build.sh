#!/bin/bash
cd `dirname $0`
sudo docker stop tods_web
sudo docker rm tods_web
sudo docker rmi tods_web
rm -rf ./dist/*

set -e
npm run build
sudo docker build -t tods_web .
