[![codecov](https://codecov.io/gh/and07/boilerplate-go/branch/master/graph/badge.svg)](https://codecov.io/gh/and07/boilerplate-go)


![](https://github.com/and07/boilerplate-go/.github/workflows/main.yml/badge.svg)


# Standard Go Project Layout (Boilerplate-GO)

This is a basic layout for Go application projects. 

[Go standards project layout](https://github.com/golang-standards/project-layout)

[Local kubernetes setup with minikube on Mac OS X](https://hackernoon.com/local-kubernetes-setup-with-minikube-on-mac-os-x-eeeb1cbdc0b)

[Using Helm to deploy to Kubernetes](https://daemonza.github.io/2017/02/20/using-helm-to-deploy-to-kubernetes/)

[Kubernetes NodePort vs LoadBalancer vs Ingress? When should I use what?](https://medium.com/google-cloud/kubernetes-nodeport-vs-loadbalancer-vs-ingress-when-should-i-use-what-922f010849e0)

https://qiita.com/sotoiwa/items/993990edf2bb98af7c1d#grafana

```sh
$ echo "$(minikube ip) prometheus.minikube" | sudo tee -a /etc/hosts 
$ echo "$(minikube ip) alertmanager.minikube" | sudo tee -a /etc/hosts 
$ echo "$(minikube ip) grafana.minikube" | sudo tee -a /etc/hosts 
$ echo "$(minikube ip) jaeger.minikube" | sudo tee -a /etc/hosts
$ echo "$(minikube ip) boi.minikube" | sudo tee -a /etc/hosts
$ echo "$(minikube ip) private-boi.minikube" | sudo tee -a /etc/hosts
```


```sh
helm package jaeger --debug
helm package boilerplate-go-chart --debug
```

```sh
$ helm install --name jaeger jaeger-0.1.0.tgz

$ helm install --name boi  boilerplate-go-chart-0.1.0.tgz 

$ helm install --name prometheus --namespace monitoring -f prometheus-values.yaml stable/prometheus

$ helm install --name grafana --namespace monitoring -f grafana-values.yaml stable/grafana
```

```sh
$ kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

```sh
helm del --purge jaeger
helm del --purge boi
helm del --purge grafana
helm del --purge prometheus
```

## GRPC

```
grpc_cli ls localhost:8842 -l 

grpc_cli call localhost:8842 HttpBodyExampleService.HelloWorld ''


grpc_cli call localhost:8842 BlockchainService.Address 'address:"Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6", height:11'

grpc_cli call localhost:8842 BlockchainService.Subscribe 'query:"testete"'

```

## GRPC http proxy
```
curl -X GET "http://localhost:8843/helloworld" -H "accept: application/json"

curl -X GET "http://localhost:8843/address/Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6?height=1" -H "accept: application/json"
```


## swagger

```
http://localhost:8888/swaggerui/
```
