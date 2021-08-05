# Bino

The booking bot at your service
### Requirements

- [Go](https://golang.org/) 1.16
- [Docker](https://www.docker.com/) 18.06.0+

### Configure Database

Both `MYSQL_ROOT_PASSWORD` and `MYSQL_DATABASE` are required to run the database. They are made available to the database via a file named `.env.docker`.

For a quick setup you can copy `.env-sample`:

```shell
cp .env-sample .env.docker
```
### Run the bino bot

#### Setup Discord Token

Bino needs to connect to a Discord server to run, so you'll need to provide a bot token. Check [WriteBots](https://writebots.com) great [tutorial](https://www.writebots.com/discord-bot-token/) to setup one.

The token needs to be available via an environment variable named `BINO_DISCORD_TOKEN`. You can do it by running:

```shell
export BINO_DISCORD_TOKEN=<your_token_goes_here>
```

#### Install dependencies

It can be done via:

```shell
go mod download
```

#### Run app

You can run the app locally via

```shell
go run main.go
```

Or using docker-compose via:

```shell
docker-compose up
```

#### Developing with Docker

You may use Docker as a development environment. For that you can run:


```shell
docker-compose run --rm app bash
```

This will open a bash session inside a container with go pre-installed and all the dependencies already downloaded. It will also mount the current directory on the container so file changes on the host are reflected on the container.

Just like on local development, the app can be started using:

```shell
go run main.go
```

### 

docker run --rm \
--name binodb \
-p 0.0.0.0:3306:3306 \
-e MYSQL_ROOT_PASSWORD="eciladabino" \
-e MYSQL_DATABASE=binodb \
-d mysql:8

mysql -h127.0.0.1 -uroot -peciladabino binodb