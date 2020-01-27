[![Build Status](https://travis-ci.com/otus-kuber-2019-12/wbe7_platform.svg?branch=master)](https://travis-ci.com/otus-kuber-2019-12/wbe7_platform)

# wbe7_platform


kubernetes-intro
-
#### Run kubernetes-intro

`kubectl apply -f kubernetes-intro/web-pod.yaml`

`kubectl apply -f kubernetes-intro/frontend-pod-healthy.yaml`

#### Expose kubernetes-intro

`kubectl port-forward web 80:8000`

`kubectl port-forward frontend 81:8080`

#### Explore kubernetes-intro

http://127.0.0.1:80

http://127.0.0.1:81

kubernetes-controllers
-
#### Run kubernetes-controllers

`kubectl apply -f kubernetes-controllers/frontend-deployment.yaml`

`kubectl apply -f kubernetes-controllers/paymentservice-deployment.yaml`

`kubectl apply -f kubernetes-controllers/node-exporter-daemonset.yaml`

#### Expose kubernetes-controllers

`kubectl port-forward frontend-* 80:8080`

`kubectl port-forward node-exporter-* 9100:9100`

#### Explore kubernetes-controllers

http://127.0.0.1:80

http://127.0.0.1:9100/metrics

kubernetes-security
-
#### Run kubernetes-security

`kubectl apply -f kubernetes-security/task01/`

`kubectl apply -f kubernetes-security/task02/`

`kubectl apply -f kubernetes-security/task03/`


#### Explore kubernetes-security

`kubectl auth can-i get deployments --as system:serviceaccount:default:bob`

yes

`kubectl auth can-i get pods --as system:serviceaccount:default:dave`

no

`kubectl auth can-i get deployments --as system:serviceaccount:prometheus:carol`

no

`kubectl auth can-i list pods --as system:serviceaccount:prometheus:carol -n prometheus`

yes

`kubectl auth can-i list pods --as system:serviceaccount:prometheus:carol`

yes

`kubectl auth can-i get deployments --as system:serviceaccount:dev:jane`

no

`kubectl auth can-i create deployments --as system:serviceaccount:dev:jane -n dev`

yes

`kubectl auth can-i get deployments --as system:serviceaccount:dev:ken`

no

`kubectl auth can-i get deployments --as system:serviceaccount:dev:ken -n dev`

yes

`kubectl auth can-i create deployments --as system:serviceaccount:dev:ken -n dev`

no

kubernetes-networks
-
#### Run kubernetes-networks

`kubectl apply -f kubernetes-networks/web-deploy.yaml`

`kubectl apply -f kubernetes-networks/web-svc-cip.yaml -f kubernetes-networks/web-svc-lb.yaml -f kubernetes-networks/web-svc-headless.yaml`

`kubectl apply -f kubernetes-networks/web-ingress.yaml`

Со *

`kubectl apply -f kubernetes-networks/coredns/`

`kubectl apply -f kubernetes-networks/dashboard/`

`kubectl apply -f kubernetes-networks/canary/app-v1.yaml -f kubernetes-networks/canary/ingress-v1.yaml`

`kubectl apply -f kubernetes-networks/canary/app-v2.yaml`

`kubectl apply -f kubernetes-networks/canary/ingress-v2-canary.yaml`

#### Expose kubernetes-networks

В зависимости от последовательности развертывания LB сервиса nginx-ingress и web-svc-lb используемые IP адреса могут отличаться. Для проверки IP адреса можно выполнить следующие команды:

`kubectl get svc`

`kubectl get svc -n ingress-nginx `

#### Explore kubernetes-networks

После того, как узнали IP адрес балансировщика ingress и сервиса web-svc-lb, проверяем наши сервисы:

http://172.17.255.*

http://172.17.255.*/web

Со *

nslookup web-svc-cip.default.svc.cluster.local 172.17.255.10

https://172.17.255.*/dashboard

while sleep 0.1; do curl 172.17.255.* -H "Host: my-app.com"; done

kubernetes-volumes
-
#### Run kubernetes-volumes

`kubectl apply -f kubernetes-volumes/`

#### Expose kubernetes-volumes

`kubectl port-forward minio-0 9000:9000`

#### Explore kubernetes-volumes

http://127.0.0.1:9000


| HW Dashboard                                                                                                         |
| :----------------------------------------------------------------------------------------------------------------------------------- |
| [**Homework**](https://github.com/orgs/otus-kuber-2019-12/projects/1)<br/> Panel for review PR (Pull Requests). |