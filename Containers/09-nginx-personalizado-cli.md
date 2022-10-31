## 09 - Nginx Personalizado usando el CLI

En este capítulo vamos a crear un contenedor de Nginx personalizado usando el CLI de Docker.
Modificaremos el fichero index.html desde el volumen y desde local para ver los cambios

### 09.1 - Crear el fichero index.html

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> mkdir -p /tmp/html
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> echo "<h1>Hola UPM</h1>" > /tmp/html/index.html
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -d -p 8080:80 -v /tmp/html:/usr/share/nginx/html:ro --name mynginx_volume nginx
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

Hola UPM
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                  NAMES
c3b0b0b5b0f1        nginx               "/docker-entrypoint.…"   2 minutes ago       Up 2 minutes
```

### 09.2 - Modificar el fichero index.html desde el volumen

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> echo "Adios UPM" > /tmp/html/index.html
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> curl localhost:8080

Adios UPM
```


