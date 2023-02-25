#!/bin/bash
docker rm tods_web
docker rmi tods_web:1.0
docker build -t tods_web:1.0 .
