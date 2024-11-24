#!/bin/bash

TAG=dev

# replace image of queue-proxy
kubectl -n knative-serving patch image queue-proxy --patch --type=merge \
    '{"spec":{"image":"docker.io/bonavadeur/nonna:'${TAG}'"}}'
kubectl -n knative-serving patch configmap config-deployment --patch \
    '{"data":{"queue-sidecar-image":"docker.io/bonavadeur/nonna:'${TAG}'"}}'
