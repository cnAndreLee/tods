#!/bin/bash
rm tods_server
go build
docker rm tods_server
docker rmi tods_server:1.0
docker build -t tods_server:1.0 .
