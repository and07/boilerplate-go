# `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).


# Kafka

## Running, this will take time to download and startup

```
docker-compose -f docker-compose-kafka.yml up
```
## You can then access various ui

```
http://localhost:8001 # kafdrop - simple admin
http://localhost:8000 # kafka-topics
```

## list topics

More commands: https://kafka.apache.org/quickstart

```sh
# list topics
docker-compose -f docker-compose-kafka.yml run cli kafka-topics.sh --list --zookeeper zookeeper:2181

# create a topic
docker-compose -f docker-compose-kafka.yml run cli kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic obb-test

# send data to kafka
docker-compose -f docker-compose-kafka.yml run cli kafka-console-producer.sh --broker-list zookeeper:9092 --topic obb-test


#
# produce and consume messages
# open 2 terminal windows
#

# producer
docker-compose -f docker-compose-kafka.yml run cli kafka-console-producer.sh --broker-list kafka:9092 --topic obb-test
# type some stuff now once it loads up


#consumer - should start seeing what you type
docker-compose -f docker-compose-kafka.yml run cli kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic obb-test
```