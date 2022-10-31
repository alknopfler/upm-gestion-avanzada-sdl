## 04 - Name and rm option to remove container

The `name` option allows you to give a name to the container. This is useful to identify the container. The `rm` option allows you to remove the container when it is stopped. This is useful to avoid having a lot of containers in your system.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run --name my_container -it --rm alpine sh
```

```shell
/ # echo "do whatever you want here inside the container"
```

```shell
/ # exit
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

