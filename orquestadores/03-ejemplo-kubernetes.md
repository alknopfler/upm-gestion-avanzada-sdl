# Ejemplo deployment con kubernetes como orquestador

En este documento vamos a desplegar un servidor y un cliente con kubernetes

## Servidor

En ese caso para el servidor vamos a utilizar:
 - Un Servicio que expone el puerto 8080
 - Un statefulset que levanta 3 replicas del servidor usando el container que hemos generado anteriormente. 
La característica de este tipo de despliegue es que se asegura que siempre hay 3 replicas del servidor levantadas, y que si alguna de ellas falla, se levanta otra. Además, gestiona una aplicacion con estado dado que tenemos un volumen que se monta en cada una de las replicas para utilizar la base de datos

```bash
kubectl apply -f orquestadores/kubernetes/service.yaml
kubectl apply -f orquestadores/kubernetes/statefulset.yaml
```

## Cliente

En este caso para el cliente vamos a utilizar un deployment

```bash
kubectl apply -f orquestadores/kubernetes/client-deployment.yaml
```

## Conclusiones y preguntas para debatir
¿Qué ocurre si se cae un pod? Si elimino el pod del servidor por ejemplo.

¿podríamos jugar para reducir la memoria de los recursos del container?

¿Y si no queremos acceso externo al servidor?

