# Ejemplo sin orquestador

En este documento vamos a desplegar un servidor y un cliente sin orquestación para entender la necesidad, así como la diferencia entre desplegarlos de forma manual y con orquestación.

### Servidor:
Lo primero que hacemos es pull de la imagen ya construida de nuestro servidor:
```bash
docker pull alknopfler/upm-master-api-servidor:latest
```

Lo siguiente que vamos a hacer es levantar el servidor, para ello debemos conocer las especificaciones:

- Puerto: 8080
- Volume: /tmp
- Type: TCP
- Daemon: yes
- Ruta de acceso Web: /api/v1/accounts

Tal y como vimos en el capítulo de containers vamos a levantar el proceso:
```bash
docker run -d -p 8080:8080 -v /tmp:/tmp --name upm-master-api-servidor alknopfler/upm-master-api-servidor:latest
```
