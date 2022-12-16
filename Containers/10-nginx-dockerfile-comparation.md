## 10 - nginx Dockerfile comparation

In this example we will compare the Dockerfile of the official nginx image with a custom nginx image build from scratch. We will see the differences between the two images.

### 10.1 - Official nginx image

The official nginx image is based on the alpine image. We can see the Dockerfile of the official nginx image in the [Docker Hub](https://hub.docker.com/_/nginx/).

```shell
echo "<h1>Prueba</h1>" > /tmp/index.html
```
```Docker
FROM nginx:alpine
COPY /tmp/index.html /usr/share/nginx/html/index.html
```
```shell
docker build -t mynginx-alpine .


docker run -it -p 8080:80  --rm --name simple-nginx-running-app mynginx-alpine
```


### 10.2 - Custom nginx image
```Docker
# Pull the minimal Ubuntu image
FROM ubuntu

# Install Nginx
RUN apt-get -y update && apt-get -y install nginx

# Expose the port for access
EXPOSE 80/tcp

# Run the Nginx server
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
```
```shell

docker build -t mynginx-complex .


docker run -d -p 8080:80 --rm --name complex-nginx-running-app mynginx-complex
```


¿Cuál es la diferencia entre las dos imágenes? ¿Cuál es más pequeña? ¿Cuál es más segura? ¿Cuál es más fácil de mantener?
