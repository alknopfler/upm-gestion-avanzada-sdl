# Serverless


En este ejemplo vamos a implementar nuestro código usando una función en lugar de desplegar de forma permanente el contenedor.
Como podremos ver, el resultado esperado es:

- Solo cuanod invoco la URL se genera un evento que permite lanzar el servidor y ponerlo a funcionar
- Mientras tanto se queda apagado (sin consumo de recursos)

Para este ejemplo, vamos a usar nuestro servidor y sin cambiar nada de su código lo vamos a convertir en función, usando como proveedor Knative en lugar de utilizar un proveedor de servicios serverless.

Para ello es imprescindible haber instalado los pre-requisitos anteriores en nuestro servidor


## Instalación de knative en kubernetes

He creado un script que instala knative en kubernetes, para ello ejecutamos:

```shell
./preparation.sh 
```

O bien podemos ejecutar los pasos a mano:

```shell
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-crds.yaml
kubectl wait --for=condition=Established --all crd

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-core.yaml
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n knative-serving > /dev/null

kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.0.0/kourier.yaml
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n kourier-system
kubectl wait pod --timeout=-1s --for=condition=Ready -l '!job-name' -n knative-serving

EXTERNAL_IP=$(kubectl -n kourier-system get service kourier -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
echo EXTERNAL_IP=$EXTERNAL_IP
KNATIVE_DOMAIN="$EXTERNAL_IP"

kubectl patch configmap -n knative-serving config-domain -p "{\"data\": {\"$KNATIVE_DOMAIN\": \"\"}}"

kubectl patch configmap/config-network \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"ingress.class":"kourier.ingress.networking.knative.dev"}}'
```

## Verificar Knative instalado y listo

```shell
kubectl get pods -n knative-serving
kubectl get pods -n kourier-system
kubectl get svc  -n kourier-system
```

## Despliegue de nuestra función usando la imagen del servidor

```shell
cat <<EOF | kubectl apply -f -
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: myserver
spec:
  template:
    spec:
      containers:
        - image: alknopfler/upm-master-api-servidor:latest
          ports:
            - containerPort: 8080
          env:
            - name: TARGET
              value: "Knative"
EOF

kubectl wait ksvc myserver --all --timeout=-1s --for=condition=Ready
SERVICE_URL=$(kubectl get ksvc myserver -o jsonpath='{.status.url}')
URL="$SERVICE_URL/api/v1/accounts"
curl $URL
```

## verificar que crea los pods / funciones 

```shell
kubectl get pod -A | grep myserver
```

Después de un tiempo el pod debería ser borrado porque ha terminado de ejecutar nuestra función:

```shell
kubectl get pod -A | grep myserver
```

Si volvemos a ejecutar el curl, el pod debería ser cread de nuevo y ejecutar la función de nuestro servidor:

```shell
curl $URL
```

```shell
kubectl get pod -A | grep myserver
```


## Limpieza del entorno

```shell
kn service delete myserver

kubectl delete -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-crds.yaml
kubectl delete -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-core.yaml
kubectl delete -f https://github.com/knative/net-kourier/releases/download/knative-v1.0.0/kourier.yaml

kubectl delete ns knative-serving
kubectl delete ns kourier-system

```