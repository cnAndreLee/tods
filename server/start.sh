#!/bin/bash

docker stop tods_server
docker rm tods_server
docker run -d --name tods_server -p 8000:8000 -e TODS_DB_HOST=10.0.0.10 -e TODS_DB_PORT=54321 -v /home/ubuntu/file:/root/file tods_server:1.0
