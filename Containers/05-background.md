## 05 - Run a container in background

The `d` option allows you to run a container in background. This is useful to run a container in background and not to block the terminal.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -d alpine sh -c "while true; do echo hello world; sleep 1; done"
```

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
f3b8d4b8d0b0        alpine              "sh -c 'while true..."   2 seconds ago       Up 1 second                             boring_bell
```
