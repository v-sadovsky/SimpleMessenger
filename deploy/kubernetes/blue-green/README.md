# Реализация стратегии blue-green деплоя на примере auth сервиса

Данный вид развертывания будем выполнять вручную на основе примитива Deployment Kubernetes.

### Реализуем маршрутизацию трафика к репликам auth сервиса через Ingress

Сперва создадим NodePort сервис (в манифесте указываем метку version: blue для auth сервиса):
```
$ kubectl apply -f deploy/kubernetes/service_node_port.yaml
service/auth-service created

# Проверяем
$ kubectl --namespace auth-service get svc
NAME           TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
auth-service   NodePort   10.102.94.229   <none>        80:30000/TCP   73s
```

Разворачиваем Ingress контроллер на базе nginx
```
# Устанавливаем ingress в minikube
$ minikube addons enable ingress

# И создаем его инстанс на основе манифест-файла
$ kubectl apply -f deploy/kubernetes/ingress.yaml 

# Проверяем
$ kubectl --namespace auth-service get ingress
NAME        CLASS   HOSTS            ADDRESS        PORTS   AGE
messenger   nginx   messenger.info   192.168.49.2   80      131m
```

### Деплоим auth сервис версии v1.0.0
   
При запуске сервиса в консоли ожидаем увидеть сообщение:
`"Starting Auth service v1.0.0 on port :80"`

Собираем образ на основе Dockerfile_auth файла:
```
$ docker build -t auth-service:v1.0.0 -f build/package/Dockerfile_auth .
``` 

Убеждаемся, что собранный образ доступен внутри миникуба:
```
$ minikube image ls --format table
|-----------------------------------------|----------|---------------|--------|
|                  Image                  |   Tag    |   Image ID    |  Size  |
|-----------------------------------------|----------|---------------|--------|
| docker.io/library/auth-service          | v1.0.0   | 6a3c938717e36 | 4.92MB |
``` 

Разворачиваем blue версию auth сервиса в неймспейсе auth-service:
```
$ kubectl apply -f deploy/kubernetes/blue-green/blue-deployment.yaml
deployment.apps/auth-service-blue created

#Проверяем, что все реплики созданы и находятся в состоянии READY и RUNNING
$ kubectl --namespace auth-service get pods -l app=auth-service,version=blue
NAME                                READY   STATUS    RESTARTS   AGE
auth-service-blue-6d96c4f499-khpn6   1/1     Running   0          3m16s
auth-service-blue-6d96c4f499-vh7s6   1/1     Running   0          3m16s
auth-service-blue-6d96c4f499-wx5sg   1/1     Running   0          3m16s
```

Пробуем извне достучаться до blue версии сервиса через ingress:
```
$ curl --resolve "messenger.info:80:$( minikube ip )" -i http://messenger.info/
HTTP/1.1 200 OK
Date: Sat, 03 Aug 2024 14:10:55 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 33
Connection: keep-alive

Welcome to Auth Service v1.0.0 
``` 

Видим ожидаемый ответ сервиса `Welcome to Auth Service v1.0.0`.

### Подготавливаем и разворачиваем обновленную (green) версию auth сервиса 

Вносим правки в main.go файл auth сервиса и собираем образ обновленной версии с меткой v2.0.0:
```
$ docker build -t auth-service:v2.0.0 -f build/package/Dockerfile_auth .
``` 

Убеждаемся, что оба образа доступен внутри миникуба:
```
$ minikube image ls --format table
|-----------------------------------------|----------|---------------|--------|
|                  Image                  |   Tag    |   Image ID    |  Size  |
|-----------------------------------------|----------|---------------|--------|
| docker.io/library/auth-service          | v1.0.0   | 6a3c938717e36 | 4.92MB |
| docker.io/library/auth-service          | v2.0.0   | 6a3c938717e36 | 4.92MB |
``` 

Разворачиваем green версию образа:
```
$ kubectl apply -f deploy/kubernetes/blue-green/green-deployment.yaml
deployment.apps/auth-service-green created
``` 

Проверяем, что все реплики созданы и находятся в состоянии READY и RUNNING:
```
$ kubectl --namespace auth-service get pods -l app=auth-service,version=green
NAME                                  READY   STATUS    RESTARTS   AGE
auth-service-green-76d7f65dcf-bxhfj   1/1     Running   0          22s
auth-service-green-76d7f65dcf-cw5zv   1/1     Running   0          22s
auth-service-green-76d7f65dcf-dz8jh   1/1     Running   0          22s
```

На данный момент в кластере развернуты обе версии auth сервиса:
```
$ kubectl --namespace auth-service get deployment
NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
auth-service-blue    3/3     3            3           53m
auth-service-green   3/3     3            3           2m4s
```

Весь трафик пока что идет в blue версию auth сервиса:
```
$ curl --resolve "messenger.info:80:$( minikube ip )" -i http://messenger.info/
HTTP/1.1 200 OK
Date: Sat, 03 Aug 2024 14:23:43 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 33
Connection: keep-alive

Welcome to Auth Service v1.0.0 
```

###  Переключаем трафик на реплики, контролируемые зеленым развертыванием

Убедившись, что новая версия auth сервиса работоспособна и готова обрабатывать запросы пользователей, можно переключить трафик с blue реплик на новые green реплики.

Для этого редактируем манифест NodePort сервиса так, чтобы его селектор меток указывал на поды со значением green меток, и применяем изменения:
```
$ kubectl apply -f deploy/kubernetes/service_node_port.yaml
service/auth-service configured
``` 

Теперь весь трафик должен роутиться на green поды:
```
$ curl --resolve "messenger.info:80:$( minikube ip )" -i http://messenger.info/
HTTP/1.1 200 OK
Date: Sat, 03 Aug 2024 14:26:23 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 33
Connection: keep-alive

Welcome to Auth Service v2.0.0
``` 

Старую blue версию деплоя можно удалить:
```
$ kubectl --namespace auth-service delete deployment auth-service-blue
deployment.apps "auth-service-blue" deleted
```