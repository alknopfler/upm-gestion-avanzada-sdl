# Ejemplo sin orquestador

En este documento vamos a desplegar un servidor y un cliente con un docker-compose que aunque no es un orquestador, si que nos sirve para poder entender conceptos importantes de orquestación.

## Run docker-compose

Para empezar, debemos acceder al directorio donde hemos clonado el repositorio del máster y hacer lo siguiente:
```bash
cd orquestadores/docker-compose
docker-compose up --build
```

## Conclusiones y preguntas para debatir

- ¿Realmente es escalable?
- ¿Qué ocurre con los recursos de la máquina?
- ¿Cómo se podría correr en diferentes máquinas?
- ¿Cómo se podría saber si un servicio se ha caído?
