// +heroku goVersion go1.14
// +heroku install ./cmd/...

module github.com/and07/boilerplate-go

go 1.14

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/golang/protobuf v1.3.5
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/jnewmano/grpc-json-proxy v0.0.0-20200227201450-d2f9a1e2ec3d
	github.com/klim0v/grpc-gateway-example v0.0.0-20200421090150-0b0c1702054d
	github.com/kshvakov/clickhouse v1.3.6
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mssola/user_agent v0.5.2
	github.com/olivere/elastic v6.2.17+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	github.com/prometheus/client_golang v1.5.1
	github.com/rakyll/statik v0.1.7
	github.com/ribice/golang-swaggerui-example v0.0.0-20180611180427-1e7622a30e50
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc
	github.com/ua-parser/uap-go v0.0.0-20200325213135-e1c09f13e2fe
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	github.com/valyala/fasthttp v1.12.0
	go.elastic.co/apm/module/apmgrpc v1.7.2
	go.uber.org/zap v1.14.1
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/appengine v1.4.0
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c
	google.golang.org/grpc v1.29.0
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	sourcegraph.com/sourcegraph/go-diff v0.0.0-20171119081133-3f415a150aec // indirect
)
