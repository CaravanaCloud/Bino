# Bino

The booking bot at your service

### Configure Database 

Create copy file

```shell
$ mv .env-sample .env.docker
```

Run the command for import dump in `mysql`

```shell
$ docker exec -i binodb mysql -uroot -pdocker bino < ./db/Bino.sql
```

### Run the bino bot

Execute this command for create container:

```shell
$ docker-compose run --service-ports --rm app bash
```

Inside container run the command for compile bino and run:

```shell
$ export BINO_DISCORD_TOKEN=''
$ go run main.go
```
