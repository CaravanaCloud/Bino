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
