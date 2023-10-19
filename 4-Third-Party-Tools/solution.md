## Solution

## Install ingress-nginx

```
cd ingress
kubectl apply -f .
```

## Install ArgoCD


```
cd argocd
kubectl apply -f .
```

## Add to etc/hosts
Open /etc/hosts with your editor and create an entry for 127.0.0.0 for argo.example.com
Ingress relies on routing by both host and by route. 

## Get password for argo

We need the initial login password for argo which is stored in a secret.
```
kubectl get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```

## Add our local cluster to argocd so it has the permissions to manage it
```
argocd login argo.example.com
argocd cluster add docker-desktop
```

## Deploy the contents of the infra folder

We have a deployment from before, but also an ingress object that directs external traffic to our internal services.
```
cd infra
kubectl apply -f .
```

## Manage changes to our infra by creating a new app in argo with the GUI or with CLI
To access the gui navigate to argo.example.com with your browser

to create the new app with CLI:
```
cd argoApps
kubectl apply -f .
```

