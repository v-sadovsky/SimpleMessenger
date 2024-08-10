# Реализация стратегии canary деплоя на примере auth сервиса

Данный вид развертывания будем выполнять вручную на основе примитивов Deployment Kubernetes.

### Деплоим базовую версию auth сервиса (v1.0.0):
```
# Собираем образ
$ docker build -t base-app:v1.0.0 -f build/package/Dockerfile_auth .

# Деплоим 1 реплику, чтобы проще было отслеживать результаты 
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/base-deployment.yaml

# Создаем сервис ClusterIP для доступа к подам базовой версии
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/base-cluster-ip.yaml
```

### Аналогичным образом деплоим новую версию auth сервиса (v2.0.0):
```
# Собираем образ
$ docker build -t canary-app:v2.0.0 -f build/package/Dockerfile_auth .

# Деплоим 1 реплику, чтобы проще было отслеживать результаты
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/canary-deployment.yaml

# Создаем сервис ClusterIP для доступа к подам базовой версии
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/canary-cluster-ip.yaml
```

### Настраиваем внешний доступ и canary паттерн деплоя через ingress
```
# Для базовой версии
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/base-ingress.yaml

# Для новой версии (перенаправление трафика на эту сборку - 30%)
$ kubectl --namespace canary-ns apply -f deploy/kubernetes/canary/canary-ingress.yaml
```

### Проверяем результат
```
# Делаем 10 запросов к апплику
$ for i in $(seq 1 10); do curl -s --resolve messenger.info:80:$( minikube ip ) messenger.info; done
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v2.0.0! 
Welcome to Auth Service v1.0.0!
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v2.0.0! 
Welcome to Auth Service v1.0.0! 
Welcome to Auth Service v2.0.0!
```







for i in $(seq 1 10); do curl -s --resolve messenger.info:80:$( minikube ip ) messenger.info; done