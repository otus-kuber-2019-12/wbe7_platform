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


| HW Dashboard                                                                                                         |
| :----------------------------------------------------------------------------------------------------------------------------------- |
| [**Homework**](https://github.com/orgs/otus-kuber-2019-12/projects/1)<br/> Panel for review PR (Pull Requests). |