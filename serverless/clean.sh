#!/bin/bash

kn service delete myserver

kubectl delete -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-crds.yaml
kubectl delete -f https://github.com/knative/serving/releases/download/knative-v1.0.0/serving-core.yaml
kubectl delete -f https://github.com/knative/net-kourier/releases/download/knative-v1.0.0/kourier.yaml

kubectl delete ns knative-serving
kubectl delete ns kourier-system


