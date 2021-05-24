# Bino

The booking bot at your service

# How to run 

## Start a PostgreSQL

```
export POSTGRES_PASSWORD=Masterkey321
export POSTGRES_DB=binodb
export POSTGRES_PORT=5444
export POSTGRES_CONTAINER=binopgsql
```

```
docker run --rm --name $POSTGRES_CONTAINER \
    -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -e POSTGRES_DB=$POSTGRES_DB \
    -p 0.0.0.0:$POSTGRES_PORT:5432 \
    -d postgres
```

```
psql -U postgres -h 127.0.0.1 -p $POSTGRES_PORT binodb -f db/Bino.sql 
```
## Run the Bot App

