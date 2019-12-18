# wbe7_platform

#### kubernetes-intro

1. Причина восстановления подов после удаления описана в PR.
2. Веб сервер написан на Golang, работает на 8000 порту. В роуте / сервер отдает файлы из директории /app. Также написан роут /_healthz для реализации healthcheck'ов в описании web-pod.yaml.
3. Написан Dockerfile для сборки приложения в контейнер. Приложение запускается от пользователя big-boss с UID 1001.
4. Написан манифест web-pod.yaml, описан init-container и liveness/readyness probe.
5. Сгенерирован манифест для разворачивания frontend микросервиса магазина hipster-shop.

*6. Исправлена причина ошибки в frontend-pod.yaml. Рабочий манифест frontend-pod-healthy.yaml.
Ошибка заключалась в том, что не были указаны переменные окружения, без которых микросервис не запускался. 


##### Запуск

Запуск web:

`kubectl apply -f web-pod.yaml`

`sudo kubectl port-forward web 80:8000`

Запуск frontend:

`kubectl apply -f frontend-pod-healthy.yaml`

`sudo kubectl port-forward web 80:8080`