# Stateless Workloads on kubernetes

## Overview
The 12-factor approach to applications prescribes the use of stateless applications. Kubernetes really shines here, as it makes declaratively deploying and scaling stateless workloads extremely easy. In this section, we will deploy a kubernetes Deployment, which is a collection of stateless Pods, and a kubernetes Service, which handles the routing to our pods within the cluster.

## Prerequisites
* Fork this repo
* Dockerhub account
* A local cluster running on your machine - Rancher Desktop is recommended

## Task
Create a Deployment manifest of our application from the previous sections with an associated Service manifest, and apply them to the local kubernetes cluster.

## Success critier
* Deployment and service running in a local cluster
* Check by port-forwarding the service to localhost.