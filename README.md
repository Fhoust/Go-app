# Go App

A small go app just in order to study Go

## Requirements

* MySQL
* Docker (If you're using container)

## Before running

This app collect some data from env, bellow you can find a list of all vars and their values:

| Variable      | Default value | Description   |
|:-------------:|:-------------:|:-------------:|
|    DB_URL     |    0.0.0.0    | Database URL  |
|    DB_USER    |     root      | Database User |
|    DB_PASS    |               | Database Pass |
|     PORT      |    3000       |   App Port    |

If the app cannot connect to database, it will exit.

### Creating the container

While the same folder of Dockerfile, run:

``` shell
docker build --tag db-client:latest .
```

## Running


``` shell
go run *.go
```

### Running as container

Starting MySQL container:

``` shell
docker run --name mysql --network host -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

Starting the app

``` shell
docker run -ti -e DB_PASS=123456 --network host db-client:latest
```

## API

### /usuarios/

| Function      | Method        | Expected         |
|:-------------:|:-------------:|:----------------:|
|  Get user     | GET           | /{id} (optional) |
|  Create new   | POST          | JSON             |
|  Update user  | UPDATE        | JSON             |
|  Delete user  | DELETE        | /{id} (optional) |

### API examples

* Usuarios

Return all users or the info of just one ID

``` shell
curl http://0.0.0.0:3000/usuarios/    ## Return all users
curl http://0.0.0.0:3000/usuarios/5   ## Return info about just ID 5
```

* Create

Create a new user in the database

``` shell
curl -d '{"nome":"Old App 2.0"}' -H "Content-Type: application/json" -X POST http://localhost:3000/usuarios/
```

* Update

This update some already existing ID

``` shell
curl -d '{"nome":"New potato"}' -H "Content-Type: application/json" -X UPDATE http://0.0.0.0:3000/usuarios/2   ## Update ID 2 nome
```

* Delete

``` shell
curl -X DELETE http://0.0.0.0:3000/usuarios/    ## Delete all users
curl -X DELETE http://0.0.0.0:3000/usuarios/5   ## Delete just ID 5
```
