// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/and07/boilerplate-go

go 1.14

require (
	github.com/ClickHouse/clickhouse-go v1.5.4
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/hashicorp/go-hclog v1.2.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/jnewmano/grpc-json-proxy v0.0.6
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.4
	github.com/mailgun/mailgun-go/v4 v4.6.1
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/markbates/goth v1.69.0
	github.com/mssola/user_agent v0.5.3
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/olivere/elastic v6.2.37+incompatible
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	github.com/prometheus/client_golang v1.12.1
	github.com/satori/go.uuid v1.2.0
	github.com/sendgrid/rest v2.6.9+incompatible // indirect
	github.com/sendgrid/sendgrid-go v3.11.1+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
	github.com/streadway/amqp v1.0.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75
	github.com/ua-parser/uap-go v0.0.0-20211112212520-00c877edfe0f
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/valyala/fasthttp v1.34.0
	go.elastic.co/apm/module/apmgrpc v1.15.0
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20220329172620-7be39ac1afc7
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/examples v0.0.0-20220329220628-b6873c006da7 // indirect
	google.golang.org/protobuf v1.28.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
