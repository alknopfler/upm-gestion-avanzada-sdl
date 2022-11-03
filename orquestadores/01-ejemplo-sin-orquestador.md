# Ejemplo sin orquestador

En este documento vamos a desplegar un servidor y un cliente sin orquestación para entender la necesidad, así como la diferencia entre desplegarlos de forma manual y con orquestación.

## Preparar el entorno

### Cliente:

```bash
docker pull alknopfler/upm-master-api-cliente:latest
```
### Servidor  
```bash
docker pull alknopfler/upm-master-api-servidor:latest
```

### Creamos la red para comunicar los contenedores

```bash
docker network create --driver bridge upm-net
```

## Desplegar el servidor

Lo que vamos a hacer es levantar el servidor, para ello debemos conocer las especificaciones:

- Puerto: 8080
- Volume: /tmp
- Type: TCP
- Daemon: yes
- Ruta de acceso Web: /api/v1/accounts

Tal y como vimos en el capítulo de containers, vamos a levantar el proceso:
```bash
docker run -d --network upm-net -p 8080:8080 -v /tmp:/tmp --name upm-master-api-servidor alknopfler/upm-master-api-servidor:latest

docker logs upm-master-api-servidor -f
```

## Desplegar el cliente
El cliente no tiene peculiaridades. Lo que hace esta imagen es:
- Crear 5 cuentas si el servidor está disponible
- Esperar unos segundos
- Consultar las cuentas creadas
- Si en algún punto da algún error hace reintentos hasta que el servidor esté disponible, pero después de unos intentos MUERE
- No es un daemon...Muere cuando lo hace bien, o cuando la política de reintentos termina.
```bash
docker run --network upm-net -e URL=http://upm-master-api-servidor:8080 --name upm-master-api-cliente alknopfler/upm-master-api-cliente:latest 
```

## Conclusiones y preguntas para debatir

- ¿Funciona bien verdad? Pues ahora vamos a invertir el orden. Despleguemos primero el cliente, y después el servidor. ¿Qué pasa? ¿Por qué?

- ¿Qué ocurre si durante el proceso de peticiones, el servidor tiene un fallo? Puedes probar a matar el servidor invocando http://<ruta del servidor>:8080/destroy
```
 docker run --rm --network upm-net --name mytest -ti alpine/curl:latest curl upm-master-api-servidor:8080/destroy
```

