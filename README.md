[![Build Status](https://travis-ci.com/otus-kuber-2019-12/wbe7_platform.svg?branch=master)](https://travis-ci.com/otus-kuber-2019-12/wbe7_platform)

# wbe7_platform

##kubernetes-intro

#####Run kubernetes-intro

`kubectl apply -f kubernetes-intro/web-pod.yaml`

`kubectl apply -f kubernetes-intro/frontend-pod-healthy.yaml`

#####Expose kubernetes-intro

`kubectl port-forward web 80:8000`

`kubectl port-forward frontend 81:8080`

#####Explore kubernetes-intro

http://127.0.0.1:80

http://127.0.0.1:81

##kubernetes-controllers

#####Run kubernetes-controllers

`kubectl apply -f kubernetes-controllers/frontend-deployment.yaml`

`kubectl apply -f kubernetes-controllers/paymentservice-deployment.yaml`

`kubectl apply -f kubernetes-controllers/node-exporter-daemonset.yaml`

#####Expose kubernetes-controllers

`kubectl port-forward frontend-* 80:8080`

`kubectl port-forward node-exporter-* 9100:9100`

#####Explore kubernetes-controllers

http://127.0.0.1:80

http://127.0.0.1:9100/metrics



| HW Dashboard                                                                                                         |
| :----------------------------------------------------------------------------------------------------------------------------------- |
| [**Homework**](https://github.com/orgs/otus-kuber-2019-12/projects/1)<br/> Панель контроля review PR (Pull Request). |