# Bino

The booking bot at your service

### Configure Database

Copy file `.env-sample` to `.env.docker`

```shell
cp .env-sample .env.docker
```
### Run the bino bot

Execute this command for create container:

```shell
docker-compose run --service-ports --rm app bash
```

Inside container run the command export `BINO_DISCORD_TOKEN` variable and run application:

```shell
export BINO_DISCORD_TOKEN=''
go run main.go
```
