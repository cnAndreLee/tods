#!/bin/bash -x

cd `dirname $0`

bash ../server/build.sh
bash ../vue/build.sh

set -e

sudo docker push andreleesss/tods_server
sudo docker push andreleesss/tods_web

