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
curl -v -X POST "http://localhost:8080/signup" -H "accept: application/json"  --data '{"email":"tete@mail.ccc","password":"xyz", "username":"TEEST"}' 
curl -v -X POST "http://localhost:8080/verify/mail" -H "accept: application/json"  --data '{"email":"tete@mail.ccc","code":"PatwtPil"}' 
curl -v -X POST "http://localhost:8080/login" -H "accept: application/json"   --data '{"email":"tete@mail.ccc","password":"xyz"}' 
curl -v -X GET "http://localhost:8080/greet" -H "accept: application/json" -H "Authorization: Bearer access_token"
curl -v -X GET "http://localhost:8080/get-password-reset-code" -H "accept: application/json" -H "Authorization: Bearer access_token"
curl -v -X POST "http://localhost:8080/verify/password-reset" -H "accept: application/json"  --data '{"email":"tete@mail.ccc","code":"fsfykSBG"}' 
curl -v -X GET "http://localhost:8080/refresh-token" -H "accept: application/json" -H "Authorization: Bearer refresh_token"
curl -v -X GET "http://localhost:8080/profile" -H "accept: application/json" -H "Authorization: Bearer access_token"



curl -X POST "http://localhost:8843/user/parameters" -H "accept: application/json"  -H "Authorization: Bearer access_token" --data '{"weight":1,"height":2,"age":2,"gender":0,"eat":1 }'
curl -X GET "http://localhost:8843/user/parameters" -H "accept: application/json"  -H "Authorization: Bearer access_token"


curl -X GET "http://localhost:8843/user/exercises/default?type=" -H "accept: application/json"  -H "Authorization:  Bearer access_token"


curl -X POST "http://localhost:8843/user/trening" -H "accept: application/json"  -H "Authorization: Bearer access_token" --data '{"name":"sssss1","exercises":[{"name":"Exercise1", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"other", "uid":"", "image":"https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-velosiped.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise2", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"arms", "uid":"", "image":"https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-skruchivaniya.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise3", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"other", "uid":"", "image":"https://filzor.ru/wp-content/uploads/2021/07/685dcbc62da32d989482066ebc60de3f-768x401.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise4", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"legs", "uid":"", "image":"https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise5", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"other", "uid":"", "image":"https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-planka-na-rukah.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise6", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"arms", "uid":"", "image":"https://cdn-st1.rtr-vesti.ru/vh/pictures/xw/317/436/2.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}, {"name":"Exercise7", "duration":"20s", "relax":"20s", "count":10, "numberOfSets":3, "numberOfRepetitions":15, "type":"legs", "uid":"", "image":"https://fitseven.ru/wp-content/uploads/2020/07/uprazhneniya-na-press-podyem-nog.jpg", "video":"", "description":"bla bla bla bla bla blabla bla bla bla bla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla blabla bla", "technique":"\n\t\tStand with your hands aside.\n\t\tPick up dumbbells.\n\t\tLift the dumbbells ahead\n\t\tLower the dumbbells", "mistake":"Don’t lift the dumbbells too high"}], "interval": "30s" }'


curl -X GET "http://localhost:8843/user/trenings" -H "accept: application/json"  -H "Authorization: Bearer access_token" 


curl -X PUT "http://localhost:8843/user/trening/status" -H "accept: application/json"  -H "Authorization: Bearer access_token" --data '{ "uid":"0dde0e5c-ce0e-4432-a076-648eea6f719e","status":"Start"}'


curl -X GET "http://localhost:8843/user/trening?uid=adc71e23-a82f-44f0-b135-3abacb9f5396" -H "accept: application/json"  -H "Authorization: Bearer access_token"



curl -X POST "http://localhost:8843/user/exercise" -H "accept: application/json"  -H "Authorization: Bearer access_token" --data '{"number_of_sets":1, "number_of_repetitions": 30, "name":"sssss1"}'
curl -X GET "http://localhost:8843/user/exercises" -H "accept: application/json"  -H "Authorization: Bearer access_token"


curl -X GET "http://localhost:8843/helloworld" -H "accept: application/json"
curl -X GET "http://localhost:8843/address/Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6?height=1" -H "accept: application/json"



curl -X POST -F "attachment=@/Users/ттт/1.png" "http://localhost:8080/v1/files"
curl -X POST -F "attachment=~/1.png" "http://localhost:8080/user/image"

curl -X POST --data-binary "@/Users/ттт/1.png" "http://localhost:8080/user/image" -H "Content-Type: application/octet-stream"

```


## swagger

```sh
http://localhost:8888/swaggerui/
```
