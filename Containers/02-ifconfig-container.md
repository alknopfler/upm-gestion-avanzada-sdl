## 02 - ifconfig container

The `ifconfig` command is used to configure network interfaces. In this example we will use the `ifconfig` command to see the network configuration of the container. For sure, we could use the command we want, but we will use `ifconfig` because it is a very common command and it is easy to understand.

```shell
┌─ alknopfler[alkmini.localdomain]:~
└──> docker run alpine ifconfig
```

```shell
Unable to find image 'alpine:latest' locally
latest: Pulling from alpine 
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:02  
          inet addr:172.17.0.2  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:3 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:274 (274.0 B)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
```