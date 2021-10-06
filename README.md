# Go App

A small go app running using gRPC

## Requirements

* MySQL
* Docker (If you're using container)

## Before running

This app collect some data from env, bellow you can find a list of all vars and their values:

| Variable      | Default value | Description   |
|:-------------:|:-------------:|:-------------:|
|    DB_URL     |    0.0.0.0    | Database URL  |
|    DB_USER    |     root      | Database User |
|    DB_PASS    |    123456     | Database Pass |
|     PORT      |     3000      |   App Port    |

If the app cannot connect to database, it will exit.

### Creating the container

While the same folder of Dockerfile, run:

``` shell
docker build --tag db-client:latest .
```

## Running


``` shell
go run server/main.go
go run client/main.go
```

### Running as container

Starting MySQL container:

``` shell
docker run --name mysql --network host -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

Starting the app

``` shell
docker run -ti --network host db-client:latest
```

## API

The app was converted to gRPC, but you can still check the old files [here](https://github.com/Fhoust/Go-app/tree/8631704338aee0b5dcd571321ab6ac4e5c03710c)

## gRPC

### Functions of server

All functions will receive and return a user struct (user.Name and User.Id)

* AddNewUser()

Insert a new user inside the database

* GetUserInfo()

Collect an id from the database

* UpdateOneUser()

Updates the value of one id

* DeleteOldUser()

Delete a user of the database

#### Examples

You will find a client example inside [client folder](https://github.com/Fhoust/Go-app/blob/master/client/main.go)