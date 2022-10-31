## 08 - Volumes in containers

The `volume` option allows you to mount a volume in the container. This is useful to share data between the host and the container.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -it --rm -v /tmp:/tmp alpine sh
```

```shell
/ # ls /tmp
```

```shell
/ # touch /tmp/test
```

```shell
/ # exit
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> ls /tmp
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES

```
```shell 
┌─ alknopfler[alkmini.localdomain]:~
└──> docker volume ls
DRIVER    VOLUME NAME
local     local_database
local     tmp
local     postgres-compose_postgres-data
local     sztp_dhcp-leases-folder

```
