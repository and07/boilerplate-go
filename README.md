[![Go Report Card](https://goreportcard.com/badge/github.com/and07/boilerplate-go)](https://goreportcard.com/report/github.com/and07/boilerplate-go)
[![codecov](https://codecov.io/gh/and07/boilerplate-go/branch/master/graph/badge.svg)](https://codecov.io/gh/and07/boilerplate-go)
[![Actions Status](https://github.com/and07/boilerplate-go/workflows/Build%20and%20Test/badge.svg)](https://github.com/and07/boilerplate-go/actions)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fand07%2Fboilerplate-go.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fand07%2Fboilerplate-go?ref=badge_shield)

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

```sh
grpc_cli ls localhost:8842 -l 

grpc_cli call localhost:8842 HttpBodyExampleService.HelloWorld ''


grpc_cli call localhost:8842 BlockchainService.Address 'address:"Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6", height:11'

grpc_cli call localhost:8842 BlockchainService.Subscribe 'query:"testete"'

```

## GRPC http proxy
```sh
curl -v -X POST "http://localhost:8080/signup" -H "accept: application/json"  --data '{"email":"and_07@mail.ru","password":"xyz"}' 
curl -v -X POST "http://localhost:8080/verify/mail" -H "accept: application/json"  --data '{"email":"and_07@mail.ru","code":"PatwtPil"}' 
curl -v -X POST "http://localhost:8080/login" -H "accept: application/json"   --data '{"email":"and_07@mail.ru","password":"xyz"}' 
curl -v -X GET "http://localhost:8080/greet" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzYTFiNDcxOC04NTkxLTQ4MjktYTNkMi0yYWM2NjZmZGM1MDYiLCJLZXlUeXBlIjoiYWNjZXNzIiwiZXhwIjoxNjQ4MzI0MzU2LCJpc3MiOiJib29raXRlLmF1dGguc2VydmljZSJ9.Q05RpWNO0hI7wIQWqlCYiQJAeQBN1cImiT0y10558nOGyST71SooIfkX2BiUPTRFSvSheQ31OYwmaQEbXdo-IhCnuG_KL7YOdyjbsks0GKO0NJ1fBcWU7LtnHDeQ1FhxcIET9529SBp30kq2wq6nBPvvtf2RNsFDK2BFHPBJXuCBji8qgQNQiTvV1y3qBEDVzhQiB27WjN7TADoYXeDJ2jgbbQ8_vPYgjBiJyffH0h1MslWYZzRlVHzoCQ0PuFfkV3wUujN4jsD9mfMEpteXzgsvcf-pNMVOLRgtxrKCJ835Ibmy-gPYKfb9clUOTunoay7sPKPJeOvmHlENZz_dpg"


curl -v -X GET "http://localhost:8080/get-password-reset-code" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzYTFiNDcxOC04NTkxLTQ4MjktYTNkMi0yYWM2NjZmZGM1MDYiLCJLZXlUeXBlIjoiYWNjZXNzIiwiZXhwIjoxNjQ4MzI0MzU2LCJpc3MiOiJib29raXRlLmF1dGguc2VydmljZSJ9.Q05RpWNO0hI7wIQWqlCYiQJAeQBN1cImiT0y10558nOGyST71SooIfkX2BiUPTRFSvSheQ31OYwmaQEbXdo-IhCnuG_KL7YOdyjbsks0GKO0NJ1fBcWU7LtnHDeQ1FhxcIET9529SBp30kq2wq6nBPvvtf2RNsFDK2BFHPBJXuCBji8qgQNQiTvV1y3qBEDVzhQiB27WjN7TADoYXeDJ2jgbbQ8_vPYgjBiJyffH0h1MslWYZzRlVHzoCQ0PuFfkV3wUujN4jsD9mfMEpteXzgsvcf-pNMVOLRgtxrKCJ835Ibmy-gPYKfb9clUOTunoay7sPKPJeOvmHlENZz_dpg"

curl -v -X POST "http://localhost:8080/verify/password-reset" -H "accept: application/json"  --data '{"email":"and_07@mail.ru","code":"fsfykSBG"}' 

curl -v -X GET "http://localhost:8080/refresh-token" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzYTFiNDcxOC04NTkxLTQ4MjktYTNkMi0yYWM2NjZmZGM1MDYiLCJDdXN0b21LZXkiOiI0ZDY2MTFkOGVlZTY5ZTY3ZWZhNDU2OWIxNTU3ODVlMzQwNTg4YzQ0ZWRiZjJmMjYwZGEzODczYTFmM2Y2ZGM3IiwiS2V5VHlwZSI6InJlZnJlc2giLCJpc3MiOiJib29raXRlLmF1dGguc2VydmljZSJ9.Y9oA9zm7m_K4ECcZz_UXEv8aTDOeHhFk1ObJjwk7XXpWlLtgZAIcLhyCXdXFMeUIgrUlHG1cPjTQ9MmiesE-yJ7GZMmvzDmF5YQBUCUHUmAu9ZqsYLkDBj_U05ePRZuyO8JuI8KTXK4p22A29ao1yxR0mZ3mQcRQ9g5KCoYsLGM7LTB_CtatwTb5aFT08BxPOkNHMUefRFFniZ-lyKm9IL4Q9QUXprX98IVORA1kFck4SZrYgp6quDvnodvy7XyNBo1QZtZidRg0jL_q__TTJBo-SZ9qwh7IfxFzZoeKpkusKUCUNouTkl51fpzsLlvWm3VNnGonNUBeYkxjZiSKMQ"



curl -v -X GET "http://localhost:8080/profile" -H "accept: application/json" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzYjg3Y2M0YS0xZDE5LTRhMzgtODc0My1jNjFkNDA1NzJlZmEiLCJLZXlUeXBlIjoiYWNjZXNzIiwiZXhwIjoxNjQ4NTQzMjIyLCJpc3MiOiJib29raXRlLmF1dGguc2VydmljZSJ9.RVE6nzgsHpWHBwJvM1LiROBexTS9umJgiQH-XsCKVWXAYLHV6g0Lowd-bMV7Di6onQPE4PalM5JXmZOc3Q0hWXRgtFFIDhgk35PomKzjHROCmd3F3Z4GDb3xpW_LCWHDPhRiLG8-ehUKoquOsTNqhSeSLghXxqMDnopWXb-huQAglg0nuCqgjzClhIsCnWXB8dLOwxMeOw_jrWRE3CP5tt2mSZ-mY4lyhGpSzs6xKCdbKLNq7AGQKDFyQRJLqpyKJKfz6bLQKU6OTEatdE9T5dX1R8DKFT0OAegowbT_v1YYcZhiARG9LZkvrUER69bbM33xxzTXFWS2KSWF7dMM2g"




curl -X GET "http://localhost:8843/helloworld" -H "accept: application/json"

curl -X GET "http://localhost:8843/address/Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6?height=1" -H "accept: application/json"
```


## swagger

```sh
http://localhost:8888/swaggerui/
```
