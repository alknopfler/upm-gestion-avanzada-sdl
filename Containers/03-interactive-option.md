## 03 - interactive option

The `interactive` option allows you to run a container in interactive mode. This means that you can interact with the container using the command line. This is useful for debugging purposes.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run -it alpine sh
```

```shell
/ # 
```

```shell
/ # ifconfig  //we could use any command we want here inside the container
```

```shell
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:02  
          inet addr:
...
```
