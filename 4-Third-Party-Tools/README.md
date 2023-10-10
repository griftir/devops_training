# Third Party Tools

## Overview
There are a number of cloud native tools that can help us in creating a 12-factor app using kubernetes. This section will be about ways we can install them.

## Prerequisites
* Docker desktop with Kubernetes enabled
* Helm CLI https://helm.sh/docs/intro/install/

## Task
* Install ingress-nginx using the yaml file
* Install ArgoCD using helm
* Create an ingress for our application
* Create an ArgoCD app to track changes to our application

## Success Criteria

* Able to access our service from outside the cluster via kuberenetes.docker.internal or whatever DNS is in /etc/hosts
* Able to make changes to manifests and have them picked up by argocd


