# 1.Database
```shell
docker run -d --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:12
# enter psql
docker exec -it postgres /bin/bash
su - postgres
psql
```

### init database
```sql
create database tods;
\c tods
insert into users (account, key, is_admin, out_date) values ('admin', '123456', true, '2099-01-01');
```

# 2.kkfileview
```
docker run -d --name kkfile -p 8012:8012 keking/kkfileview
```

# 3.TODS_WEB
Enter 'vue' and exec build.sh
```shell
docker run -d --name tods_web -p 80:80 tods_web
```


# 4.TODS_SERVER
### AT develop machine
Enter 'server' and exec build.sh

```shell
docker run -d --name tods_server -p 8000:8000 -e TODS_DB_HOST=10.0.0.10 -e TODS_DB_PORT=5432 -v /home/ubuntu/file:/root/file tods_server
```