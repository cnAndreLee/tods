# What is "TODS"
Teaching-resources on demand System


## Install by docker

```shell
# server
sudo docker run -d \
	--restart=always \
	--name tods_server \
	-p 8000:8000 \
	-e TODS_DB_HOST=10.0.0.10 \
	-e TODS_DB_PORT=5432 \
	-v /home/ubuntu/file:/root/file \
	andreleesss/tods_server

# web
sudo docker run -d \
	--restart=always \
	--name tods_web \
	-p 80:80 \
	andreleesss/tods_web

# sql
docker run -d \
	--restart=always \
	--name postgres \
	-e POSTGRES_PASSWORD=postgres \
	-p 5432:5432 \
	postgres:12

# kkfile
docker run -d \
	--restart=always \
	--name kkfile \
	-p 8012:8012 \
	keking/kkfileview

```


