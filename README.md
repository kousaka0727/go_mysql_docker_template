## go_mysql_docker_template
Docker-compose template for Golang and MySQL.
## Overview
Provides a simple environment using golang and Mysql
In addition, this repository is automatically compiled by realize
## Usage
* set .env
You can set any value you want.

``` bash 
$ cp .env.example .env
$ vi .env
```

* build and create

```
$ docker compose up -d
```

* check program execution

```
$ docker ps
1bb0633bdbc6   go_mysql_docker_template_app     "realize start"          18 minutes ago   Up 10 minutes   0.0.0.0:8000->8000/tcp              app
~~~~~~~~~~~~   go_mysql_docker_template_mysql   "docker-entrypoint.s…"   18 minutes ago   Up 10 minutes   0.0.0.0:3306->3306/tcp, 33060/tcp   mysql
~~~~~~~~~~~~   redis:6.2-buster                 "docker-entrypoint.s…"   18 minutes ago   Up 10 minutes   0.0.0.0:6379->6379/tcp              redis

$ docker logs -f 1bb0633bdbc6
[06:17:50][APP] : Watching 1 file/s 3 folder/s
[06:17:50][APP] : Install started
[06:17:51][APP] : Install completed in 1.145 s
[06:17:51][APP] : Running..
[06:18:02][APP] : hello world !! ...........
[06:18:02][APP] : Successful connection to DB !!
```