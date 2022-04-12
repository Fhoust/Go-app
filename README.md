# Go App

A small go app just in order to study Go

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
|     PORT      |     8080      |   App Port    |

If the app cannot connect to database, it will exit.

### Creating the container

While the same folder of Dockerfile, run:

``` shell
docker build --tag go-app:latest .
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
docker run -ti -e DB_PASS=123456 --network host go-app:latest
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

#### Building app container

We will need to build our container inside minikube docker in order for it be able to find our image, so run:

``` shell
$ eval $(minikube docker-env)
$ docker build --tag go-app:latest .
```

#### Deploying and testing

Inside server folder you will find a deployment file, apply to our minikube:

``` shell
$ kubectl apply -f appDeployment.yaml
```

This file will create our deployment and a service for the app, to check if everything is running, type:

``` shell
$ kubectl get po
> NAME                             READY   STATUS    RESTARTS   AGE
> go-app-5db55c49d4-kvvkk   1/1     Running   0          5s
```

Then check the service

``` shell
$ kubectl get svc
> NAME                TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
> go-app-svc   LoadBalancer   10.101.18.85   <pending>     8080:32181/TCP   10s
```

We will need an external IP to connect to our service, but as you can see our service don't have an external IP. \
How we are using a local minikube, we don't have a LoadBalancer service, as in AWS or GCP, so in order to resolve this \
we are going to open a tunnel between minikube and our local machine

``` shell
$ minikube tunnel
> Status:	
>     machine: minikube
> 	  pid: 172079
> 	  route: 10.96.0.0/12 -> 192.168.49.2
> 	  minikube: Running
> 	  services: [go-app-server-svc]
>   errors: 
> 		minikube: no errors
> 		router: no errors
> 		loadbalancer emulator: no errors


$ kubectl get svc
> NAME         TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)          AGE
> go-app-svc   LoadBalancer   10.100.145.192   10.100.145.192   8080:32674/TCP   4m7s

``` 

Use this EXTERNAL-IP in order to access your app and have fun!

## API

### /users/

| Function      | Method | Expected         |
|:-------------:|:------:|:----------------:|
|  Get user     |  GET   | /{id} (optional) |
|  Create new   |  POST  | JSON             |
|  Update user  |  PUT   | JSON             |
|  Delete user  | DELETE | /{id} (optional) |

### API examples

Remember to change **10.100.145.192** to your EXTERNAL-IP.

* Users

Return all users or the info of just one ID

``` shell
curl -d '{"name":"Old Potato"}' -H "Content-Type: application/json" -X POST http://10.100.145.192:8080/users ## Create new user
curl http://10.100.145.192:8080/users/                                                                         ## Return all users
```

* Create

Create a new user in the database

``` shell
curl -d '{"name":"Old Potato"}' -H "Content-Type: application/json" -X POST http://10.100.145.192:8080/users/
```

* Get

Create a new user in the database

``` shell
curl http://10.100.145.192:8080/users   ## Return all users
curl http://10.100.145.192:8080/users/1 ## Return only the user with ID 1
```

* Update

This update some already existing ID

``` shell
curl -d '{"name":"New potato"}' -H "Content-Type: application/json" -X PUT http://10.100.145.192:8080/users/1   ## Update user with id 1
```

* Delete

``` shell
curl -X DELETE http://10.100.145.192:8080/users/5   ## Delete just ID 5
```

# Troubleshooting

### Failed to connect

``` shell
curl http://localhost:8080/users
curl: (7) Failed to connect to localhost port 8080: Connection refused
```

Remember to use the EXTERNAL-IP in case you are using k8s.

### ErrImageNeverPull

``` shell
$ kubectl get po
> NAME                             READY   STATUS              RESTARTS   AGE
> go-app-5db55c49d4-ljzrn   0/1     ErrImageNeverPull   0          3s
```

This happens when minikube can't find our app image, make sure to follow all commands in [Building app container](
https://github.com/Fhoust/Go-app#building-app-container)

### Error or CrashLoopBackOff

``` shell
$ kubectl get po
> NAME                             READY   STATUS   RESTARTS     AGE
> go-app-5db55c49d4-8rcxr   0/1     Error    1 (2s ago)   3s
```

Or

``` shell
$ kubectl get po
> NAME                             READY   STATUS             RESTARTS     AGE
> go-app-5db55c49d4-8rcxr   0/1     CrashLoopBackOff   1 (5s ago)   7s
```

Usually this happens when our app can't connect to the database, check if MySQL container is running in your machine \
docker.

``` shell
$ eval $(minikube docker-env)

$ docker ps
CONTAINER ID   IMAGE                                 COMMAND                  CREATED       STATUS       PORTS          NAMES
5012b0043988   gcr.io/k8s-minikube/kicbase:v0.0.27   "/usr/local/bin/entr…"   6 hours ago   Up 6 hours   127.0.0.14...  minikube
```

You can check more details in our logs:

``` shell
$ kubectl logs -f -l service=server
> 2021/10/18 15:55:14 Collecting env vars
> 2021/10/18 15:55:14 Undeclared DB_USER, using default...
> 2021/10/18 15:55:14 Undeclared DB_PASS, using default...
> 2021/10/18 15:55:14 DB INFO -> URL: host.minikube.internal | User: root | Port: 8080
> 2021/10/18 15:55:14 Not able to connected to the database: dial tcp 192.168.49.1:3306: connect: connection refused
```

To start MySQL, run:

``` shell
$ docker run --name mysql --network host -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

$ docker ps
CONTAINER ID   IMAGE                                 COMMAND                  CREATED       STATUS       PORTS          NAMES
c8a302c44342   mysql:latest                          "docker-entrypoint.s…"   4 hours ago   Up 1 second                 mysql
5012b0043988   gcr.io/k8s-minikube/kicbase:v0.0.27   "/usr/local/bin/entr…"   6 hours ago   Up 6 hours   127.0.0.14...  minikube
```

## gRPC

There is an old version of this app made using gRPC, I keep as other branch.

## TODO

* Add tests