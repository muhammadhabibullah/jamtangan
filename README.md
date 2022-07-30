# Jam Tangan Service

## Setup

1. Setup config file on the [config](config) directory.

```shell
$ make config
```

2. Run docker compose to create/run the database

```shell
$ make dependency
```

3. Run migrate function to migrate initial database schema

```shell
$ make migrate
```

4. Run HTTP serve

```shell
$ make serve
```

## API

Access OpenAPI after running the HTTP service at [localhost:8000/swagger/](http://localhost:8000/swagger/)
