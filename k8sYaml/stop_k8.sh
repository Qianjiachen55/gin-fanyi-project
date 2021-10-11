#!/bin/bash

kubectl delete svc kube-go &&
kubectl delete svc kube-db-mysql &&
kubectl delete svc kube-db-redis &&
kubectl delete deployment kube-go &&
kubectl delete deployment kube-db-mysql &&
kubectl delete deployment kube-db-redis