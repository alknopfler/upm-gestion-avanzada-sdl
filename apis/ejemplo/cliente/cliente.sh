#!/bin/bash
#
# Path: apis/ejemplo/cliente/cliente.sh

# Variables
# input param usage:  ./cliente.sh <url>
URL=$1

function request() {
  for i in {1..5}; do
     curl --location --request POST "$URL/api/v1/account" \
     --header 'Content-Type: application/json' \
     --data-raw '{
         "name": "prueba",
         "balance": 5,
         "currency": "euro",
         "owner": "federico"
     }'
     if [ $? -ne 0 ]; then
       echo "[ERROR] Error waiting 30 seconds"
       sleep 15  # there was an error
     else
       echo "[DONE] Success created accounts"
       sleep 4
     fi
   done
}

function retrieve() {
  curl --location --request GET "$URL/api/v1/accounts" \
   --data-raw '' | jq
}

echo "[INFO] Client example:"
echo "[INFO] Let's create some accounts in our bank: "
request

echo "[DONE] Creation step"

echo "[INFO] Ready to start the next step to review all new accounts"
sleep 1
retrieve

echo
echo "[DONE] Review step success"
sleep 1
exit 0