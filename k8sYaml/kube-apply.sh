#!/bin/bash

kubectl apply -f deploy-mysql.yaml &&
kubectl apply -f svc-mysql.yaml &&
kubectl apply -f deploy-redis.yaml &&
kubectl apply -f svc-redis.yaml &&
kubectl apply -f deploy-go.yaml &&
kubectl apply -f svc-go.yaml &&
kubectl get svc

