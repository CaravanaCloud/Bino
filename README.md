# Bino

The booking bot at your service

### Requirements

- [Go](https://golang.org/) 1.16
- [Docker](https://www.docker.com/)

### Configure Database

Both `MYSQL_ROOT_PASSWORD` and `MYSQL_DATABASE` are required to run the database. They are made available to the database via a file named `.env.docker`.

For a quick setup you can copy `.env-sample`:

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
