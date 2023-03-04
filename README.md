# What is "TODS"
Teaching-resources on demand System


## Install by docker

#### prepare workdir
```shell
sudo mkdir -p /var/lib/tods/data
sudo mkdir -p /var/lib/tods/file
```
#### create nginx.conf
```shell
touch /var/lib/tods/nginx.conf
```

#### nginx.conf : you can instead of ip
```nginx.conf
user www-data;
worker_processes auto;
pid /run/nginx.pid;

events {
	worker_connections 768;
	# multi_accept on;
}

http {

	server{
		
		client_max_body_size 500m;
		client_body_temp_path /tmp 1 2;

		location / {
			proxy_pass http://10.0.0.10:81;
		}
		
		location ^~ /api/ {
			proxy_pass http://10.0.0.10:8000;
		}

		location ^~ /onlinePreview {
			proxy_pass http://10.0.0.10:8012;
		}
		location ^~ /pptx {
			proxy_pass http://10.0.0.10:8012;
		}
		
		
	}

}
```


#### docker
```shell
# server
sudo docker run -d \
	--restart=always \
	--name tods_server \
	-p 8000:8000 \
	-e TODS_DB_HOST=192.168.10.10 \
	-e TODS_DB_PORT=5432 \
	-e TODS_DB_USER=postgres \
	-e TODS_DB_PASSWORD=postgres \
	-e TODS_DB_DBNAME=tods \
	-v /var/lib/tods/file:/root/file \
	andreleesss/tods_server

# web
sudo docker run -d \
	--restart=always \
	--name tods_web \
	-p 81:80 \
	andreleesss/tods_web

# kkfile
sudo docker run -d \
	--restart=always \
	--name kkfile \
	-p 8012:8012 \
	keking/kkfileview

```

```shell
# sql
sudo docker run -d \
	--restart=always \
	--name postgres \
	-e POSTGRES_PASSWORD=postgres \
	-p 5432:5432 \
	-v /var/lib/tods/data:/var/lib/postgresql/data \
	postgres:12

# Create database tods
sudo docker exec -it  postgres /bin/bash
su - postgres
psql
```

```sql
create database tods;
\c tods
insert into users (account, key, is_admin, out_date) values ('admin', '123456', true, '2099-01-01');
```


```shell
sudo docker run -d \
	--name tods_nginx \
	--restart=always \
	-p 80:80 \
	-v /var/lib/tods/nginx.conf:/etc/nginx/nginx.conf:ro \
	nginx

```
