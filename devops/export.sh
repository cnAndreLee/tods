#!/bin/bash -x

cd `dirname $0`

rm -rf /tmp/tods
rm -rf /tmp/tods.tar

set -e
mkdir /tmp/tods
cp ../README.md /tmp/tods/
sudo docker save -o /tmp/tods/server.tar andreleesss/tods_server
sudo docker save -o /tmp/tods/web.tar andreleesss/tods_web

cd /tmp
sudo tar -zcvf tods.tar.gz tods

sudo rm -rf /tmp/tods

echo "Success export ! /tmp/tods.tar.gz"
