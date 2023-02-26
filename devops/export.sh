#!/bin/bash -x

cd `dirname $0`

bash ../server/build.sh
bash ../vue/build.sh

rm -rf /tmp/tods
rm -rf /tmp/tods.tar

set -e
mkdir /tmp/tods
cp ./init.md /tmp/tods/
sudo docker save -o /tmp/tods/server.tar tods_server
sudo docker save -o /tmp/tods/web.tar tods_web

cd /tmp
sudo sudo tar -zcvf tods.tar.gz tods

sudo rm -rf /tmp/tods