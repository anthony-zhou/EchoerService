# EchoerService
Echoer Service in Golang

## Run from Docker

To use the image from Docker Hub, run:

```
docker pull az42/docker-echo
```

Then run the server as follows:

```
docker run --publish 8080:8080 az42/docker-echo
```

To test, open a separate terminal window, and you should get the following behavior:

```
$ curl localhost:8080/ping

-> pong

$ curl localhost:8080/pong

-> ping
```


To build the docker image locally, run:

```
docker build --tag docker-echo .
```

(replacing `docker-echo` with whatever you'd like the image to be called).

Then you can run the container locally:

```
docker run --publish 8080:8080 docker-echo
```


