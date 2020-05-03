# note: call scripts from /scripts
SERVICE_NAME=boilerplate-go
USER=and07
VERSION=latest
REPOSITORY="$(USER)/$(SERVICE_NAME)"
QUAYVERSION = "${REPOSITORY}:$(VERSION)"

TAG=${CI_BUILD_REF_NAME}_${CI_BUILD_REF}
CONTAINER_IMAGE=docker.io/${REPOSITORY}:${TAG}
CONTAINER_IMAGE_LATEST=${QUAYVERSION}
RELEASE?=0.0.1

LOCAL_BIN:=$(CURDIR)/bin

# Check global GOLANGCI-LINT
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.13.1

# Check local bin version
ifneq ($(wildcard $(GOLANGCI_BIN)),)
GOLANGCI_BIN_VERSION:=$(shell $(GOLANGCI_BIN) --version)
ifneq ($(GOLANGCI_BIN_VERSION),)
GOLANGCI_BIN_VERSION_SHORT:=$(shell echo "$(GOLANGCI_BIN_VERSION)"|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_BIN_VERSION_SHORT:=0
endif
ifneq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_BIN_VERSION_SHORT)))"
GOLANGCI_BIN:=
endif
endif

# Check global bin version
ifneq (, $(shell which golangci-lint))
GOLANGCI_VERSION:=$(shell golangci-lint --version 2> /dev/null )
ifneq ($(GOLANGCI_VERSION),)
GOLANGCI_VERSION_SHORT:=$(shell echo "$(GOLANGCI_VERSION)"|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_VERSION_SHORT:=0
endif
ifeq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_VERSION_SHORT)))"
GOLANGCI_BIN:=$(shell which golangci-lint)
endif
endif

export GO111MODULE=on

SHELL=/bin/bash -o pipefail


APP?=$(SERVICE_NAME)
PROJECT?=gitlab.com/$(USER)/$(SERVICE_NAME)

LDFLAGS:=-X '${PROJECT}/version.Name=$(SERVICE_NAME)'\
         -X '${PROJECT}/version.ProjectID=$(CI_PROJECT_ID)'\
         -X '${PROJECT}/version.Version=$(APP_VERSION)'\
         -X '${PROJECT}/version.GoVersion=$(GO_VERSION_SHORT)'\
         -X '${PROJECT}/version.BuildDate=$(BUILD_TS)'\
         -X '${PROJECT}/version.GitLog=$(GIT_LOG)'\
         -X '${PROJECT}/version.GitHash=$(GIT_HASH)'\
         -X '${PROJECT}/version.GitBranch=$(GIT_BRANCH)'\
         -X '${PROJECT}/version.publicPortDefault='\
         -X '${PROJECT}/version.adminPortDefault='\
         -X '${PROJECT}/version.grpcPortDefault='


PKGMAP:=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,$\
        Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types



-include .env

test-env:
	@echo ${MY_ENV}

BUILD_ENVPARMS:=CGO_ENABLED=0 PORT=${PORT} PORT_DEBUG=${PORT_DEBUG} PORT_GRPC=${PORT_GRPC}
BIN?=./bin/${APP}

CONTAINER_IMAGE?=docker.io/webdeva/${APP}

.PHONY: .deps
.deps:
	$(info #Install dependencies...)
	go mod download

# install project dependencies
.PHONY: deps
deps: .deps

.PHONY: .test
.test:
	$(info #Running tests...)
	go test ./...

# run unit tests
.PHONY: test
test: .test

.PHONY: install-protoc
install-protoc:
	go mod tidy
	GOBIN=$(LOCAL_BIN)  go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

gen-protoc:
	mkdir -p "./api/${SERVICE_NAME}"
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true:./api/$(SERVICE_NAME) \
		--swagger_out=allow_merge=true,merge_file_name=api:./api/$(SERVICE_NAME) \
		--go_out=plugins=grpc:./api/$(SERVICE_NAME) ./api/*.proto

# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info #Downloading golangci-lint v$(GOLANGCI_TAG))
	go get -d github.com/golangci/golangci-lint@v$(GOLANGCI_TAG)
	go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go mod tidy
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

# run diff lint like in pipeline
.PHONY: .lint
.lint: install-lint
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.pipeline.yaml ./...

# golangci-lint diff master
.PHONY: lint
lint: .lint

# run full lint like in pipeline
.PHONY: lint-full
lint-full: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...


.PHONY: .build
.build:
	$(info #Building...)
	$(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/${APP}

# build app
.PHONY: build
build: .build

.PHONY: .run
.run:
	$(info #Running...)
	$(BUILD_ENVPARMS) go run -ldflags "$(LDFLAGS)" ./cmd/${APP}
# run app
.PHONY: run
run: .run


.PHONY: docker-kill
docker-kill:
	docker-compose -f $${DC_FILE:-deployments/docker-compose.yml} kill
	docker-compose -f $${DC_FILE:-deployments/docker-compose.yml} rm -f
	docker network rm network-$${CI_JOB_ID:-local} || true

.PHONY: docker-build
docker-build: docker-kill
	env GOBUILD=env GOOS=linux GOARCH=amd64 $(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/${APP}
	#docker build -f docker/pg-migrations.dockerfile -t pg-migrations-$${CI_JOB_ID:-local} .
	docker build --no-cache -t $(QUAYVERSION) -f build/package/project.dockerfile .
	#docker build --no-cache -f docker/itest.dockerfile -t itest-$${CI_JOB_ID:-local} .

.PHONY: docker-up
docker-up: docker-build
	docker network create network-$${CI_JOB_ID:-local}
	docker-compose -f $${DC_FILE:-deployments/docker-compose.yml} up --force-recreate --renew-anon-volumes -d

.PHONY: docker-logs
docker-logs:
	mkdir -p ./logs || true
	#docker logs postgres-$${CI_JOB_ID:-local} >& logs/postgres.log
	#docker logs pg-migrations-$${CI_JOB_ID:-local} >& logs/pg-migrations.log
	docker logs $${SERVICE_NAME}-$${CI_JOB_ID:-local} >& logs/$${SERVICE_NAME}.log
	docker logs redis-$${CI_JOB_ID:-local} >& logs/redis.log
	docker logs rabbitmq-$${CI_JOB_ID:-local} >& logs/rabbitmq.log
	docker logs elasticsearch-$${CI_JOB_ID:-local} >& logs/elasticsearch.log
	docker logs kibana-$${CI_JOB_ID:-local} >& logs/kibana.log
	docker logs prometheus-$${CI_JOB_ID:-local} >& logs/prometheus.log
	docker logs clickhouse-$${CI_JOB_ID:-local} >& logs/clickhouse.log


BEFORE_DISK_FREE=$$(df -h /)
.PHONY: docker-clean
docker-clean:
	@echo Останавливаем все контейнеры
	docker kill $$(docker ps -q) || true
	@echo Очистка докер контейнеров
	docker rm -f $$(docker ps -a -f status=exited -q) || true
	@echo Очистка dangling образов
	docker rmi -f $$(docker images -f "dangling=true" -q) || true
	@echo Очистка $${SERVICE_NAME} образов
	docker rmi -f $$(docker images --filter=reference='$${SERVICE_NAME}*' -q) || true
	#@echo Очистка itest образов
	#docker rmi -f $$(docker images --filter=reference='itest*' -q) || true
	#@echo Очистка pg-migrations образов
	#docker rmi -f $$(docker images --filter=reference='pg-migrations*' -q) || true
	@echo Очистка volume
	docker volume rm -f $$(docker volume ls -q) || true
	@echo Очистка сетей
	docker network prune -f || true
	@echo "Занятость диска до очистки:"
	@echo "${BEFORE_DISK_FREE}"
	@echo "Занятость диска после очистки:"
	@echo "$$(df -h /)"

docker-container: docker-build
    #docker build -t $(CONTAINER_IMAGE):$(RELEASE) .
	echo $(VERSION)
	#docker build --no-cache -t $(SERVICE_NAME) -f build/package/project.dockerfile .
	#docker tag $(QUAYVERSION)

docker-push: docker-container
    #docker push $(CONTAINER_IMAGE):$(RELEASE)
	docker push $(QUAYVERSION)

.PHONY: docker-push-ci
docker-push-ci:
	env GOBUILD=env GOOS=linux GOARCH=amd64 $(BUILD_ENVPARMS) go build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/${APP}
	docker build --no-cache -t ${CONTAINER_IMAGE} -f build/package/project.dockerfile .
	echo $(VERSION)
	docker tag ${CONTAINER_IMAGE} ${CONTAINER_IMAGE_LATEST}
	docker push ${CONTAINER_IMAGE}
	docker push $(CONTAINER_IMAGE_LATEST)

minikube:
	minikube delete
	minikube start --vm-driver=hyperkit
	minikube addons enable ingress

	kubectl config get-contexts
	kubectl config use-context minikube

	helm init

	minikube service list