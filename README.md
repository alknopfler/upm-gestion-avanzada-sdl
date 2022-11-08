# upm-gestion-avanzada-sdl (Ejemplos-cloud-computing Master en Gestión Avanzada en SdI)

Estos serán algunos de los ejemplos que usaré en la asignatura de Cloud Computing del Máster en Gestión Avanzada en Sdl de la Universidad Politécnica de Madrid.

# Requisitos previos

Como software vamos a necesitar:
- git
- Docker
- Docker-compose
- Kubernetes

Para evitar problemas de compatibilidad durante el curso, vamos a usar “docker-desktop” que es multi-plataforma (Windows, MacOS y Linux) y así podemos disponer de docker, docker-compose y kubernetes en local de una forma rápida y sencilla.

### Instalación de docker-desktop

- MacOS: [https://docs.docker.com/desktop/install/mac-install/](https://docs.docker.com/desktop/install/mac-install/)
- Windows: [https://docs.docker.com/desktop/install/windows-install/](https://docs.docker.com/desktop/install/windows-install/)
- Linux: [https://docs.docker.com/desktop/install/linux-install/](https://docs.docker.com/desktop/install/linux-install/)

Para obtener acceso al código de ejemplo lo primero será clonar el repositorio:

```shell
git clone https://github.com/alknopfler/upm-gestion-avanzada-sdl.git
```

# Indice de Ejemplos

## Containers

[01 - hello-world](Containers/01-hello-world.md)

[02 - ifconfig container](Containers/02-ifconfig-container.md)

[03 - interactive option](Containers/03-interactive-option.md)

[04 - name and rm option](Containers/04-name-and-rm.md)

[05 - Background option](Containers/05-background.md)

[06 - Expose Ports option](Containers/06-expose-ports.md)

[07 - Environment options](Containers/07-environment.md)

[08 - Volumes](Containers/08-volume.md)

[09 - Nginx Personalizado desde CLI](Containers/09-nginx-personalizado-cli.md)

[10 - Nginx Personalizado desde Dockerfile](Containers/10-nginx-dockerfile-comparation.md)



## Orquestadores

[01 - Ejemplo sin Orquestador](orquestadores/01-ejemplo-sin-orquestador.md)

[02 - Ejemplo con Docker Compose](orquestadores/02-ejemplo-docker-compose.md)

[03 - Ejemplo con Kubernetes](orquestadores/03-ejemplo-kubernetes.md)


## APIs 

Para obtener acceso al código lo primero será clonar el repositorio y acceder a la ruta (apis/ejemplo/servidor):

```shell
git clone https://github.com/alknopfler/upm-gestion-avanzada-sdl.git

cd upm-gestion-avanzada-sdl/apis/ejemplo/servidor
```

### Ejemplo de Swagger - Servidor con API Rest 

[01 - Swagger](apis/ejemplo/servidor/swaggerui/index.html)

### Ejemplo de router - Servidor con API Rest 

[02 - Router API - Download File](apis/ejemplo/servidor/main.go)   |   [02 - Router API - View File](https://github.com/alknopfler/upm-gestion-avanzada-sdl/blob/main/apis/ejemplo/servidor/main.go)


[03 - Modelo de Datos Api Rest - Download File](apis/ejemplo/servidor/data_model.go)   |   [03 - Modelo de Datos Api Rest - View File](https://github.com/alknopfler/upm-gestion-avanzada-sdl/blob/main/apis/ejemplo/servidor/data_model.go)

[04 - Implementación de handlers - Download File](apis/ejemplo/servidor/api.go)   |   [04 - Implementación de handlers - View File](https://github.com/alknopfler/upm-gestion-avanzada-sdl/blob/main/apis/ejemplo/servidor/api.go)


## Serverless,  IoT, Edge computing

[01 - Ejemplo Serverless](serverless/ejemplo.md)
