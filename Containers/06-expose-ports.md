## 06 - Expose Ports in Containers

The `p` option allows you to expose a port in the container. This is useful to expose a port in the container and be able to access to the service running in the container from outside the container.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -d -p 8080:80 nginx
```

```shell
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
a3ed95caeb02: Pull complete
c2e6b517cf30: Pull complete
...
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> curl localhost:8080
```

```shell
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
...
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                  NAMES   
c3b0b0b5b0f1        nginx               "/docker-entrypoint.…"   2 minutes ago       Up 2 minutes
```
