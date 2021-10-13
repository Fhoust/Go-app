# Go App

A small go app running using gRPC

## Requirements

### Standalone

* [MySQL](https://www.mysql.com/)

### Running a container

* [Docker](https://docs.docker.com/get-docker/)
* [MySQL image](https://hub.docker.com/_/mysql)

### Running in K8s 

You will need the requirements of [Docker](https://github.com/Fhoust/Go-app/#running-as-container) plus the list bellow

* [Minikube](https://minikube.sigs.k8s.io/docs/start/)
* [Kubectl](https://kubernetes.io/docs/tasks/tools/)
* [Kubectx](https://github.com/ahmetb/kubectx)
* [Kubens](https://github.com/ahmetb/kubectx) 

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

In root folder of the project, run:

``` shell
$ docker build --tag go-app-server:latest .
```

## Running


``` shell
$ go run server/main.go
$ go run client/main.go
```

### Running as container

Starting MySQL container:

``` shell
$ docker run --name mysql --network host -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

Starting the app

``` shell
$ docker run -ti --network host go-app-server:latest
```

### Running in K8s

Before starting, be sure to fulfil all requirements of [Running in K8s](https://github.com/Fhoust/Go-app/#running-in-k8s)

#### Starting and setting minikube

``` shell
$ minikube start
```

Wait for our cluster be ready, then:

``` shell
$ kubectl create namespace development
$ kubens development
```

#### Starting MySQL DB

``` shell
$ docker run --name mysql --network host -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

#### Starting minikube and building app container

We will need to build our container inside minikube docker in order for it be able to find our image, so run:

``` shell
$ eval $(minikube docker-env)
$ docker build --tag go-app-server:latest .
```

#### Deploying and testing

Inside server folder you will find a deployment file, apply to our minikube:

``` shell
$ kubectl apply -f server/serverDeployment.yaml
```

This file will create our deployment and a service for the app, to check if everything is running, type:

``` shell
$ kubectl get po
> NAME                             READY   STATUS    RESTARTS   AGE
> go-app-server-5db55c49d4-kvvkk   1/1     Running   0          5s
```

Then check the service 

``` shell
$ kubectl get svc
> NAME                TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
> go-app-server-svc   LoadBalancer   10.101.18.85   <pending>     3000:32181/TCP   10s
```

We will need an external IP to connect to our service, but as you can see our service don't have an external IP. \
How we are using a local minikube, we don't have a LoadBalancer service, as in AWS or GCP, so in order to resolve this \
we are going to open a tunnel between minikube and our local machine

``` shell
$ minikube tunnel

$ kubectl get svc
> NAME                TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
> go-app-server-svc   LoadBalancer   10.101.18.85   127.0.0.1     3000:32181/TCP   20s
``` 

Now we are ready, just run the client app:

``` shell
$ go run client/main.go Potato
> 2021/10/13 17:04:47 Added Potato
> 2021/10/13 17:04:47 The id of Potato is 15
> 2021/10/13 17:04:47 Updating 15 to New Potato
> 2021/10/13 17:04:47 In database 15 has the name New Potato
> 2021/10/13 17:04:47 15 was deleted
> 2021/10/13 17:04:47 Now database has this value for 15:
```

## API REST

The app was converted to gRPC, but you can still check the old files [here](
https://github.com/Fhoust/Go-app/tree/8631704338aee0b5dcd571321ab6ac4e5c03710c)

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