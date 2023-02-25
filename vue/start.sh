#!/bin/bash

docker stop tods_web
docker rm tods_web
docker run -d --name tods_web -p 80:80 tods_web:1.0
