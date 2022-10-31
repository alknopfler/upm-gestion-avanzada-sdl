## 07 - Environment Variables

Using -e or --env you can set environment variables inside the container. This is useful for setting up the environment for the application you are running inside the container.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -it --rm -e MY_ENV_VAR=hello alpine sh
```

```shell
/ # echo $MY_ENV_VAR

hello

/ # exit
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```
