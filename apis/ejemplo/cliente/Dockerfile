FROM alpine/curl:latest
RUN apk add --no-cache bash jq
ENV URL="http://localhost:8080"
COPY ./cliente.sh /
RUN chmod +x /cliente.sh
#The argument passed in cmd is the server address
CMD [ "sh", "-c", "/cliente.sh $URL" ]
