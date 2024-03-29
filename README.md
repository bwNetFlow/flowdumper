# bwNetFlow Example: Consumer Dumper

This bwNetFlow Kafka Consumer reads flows from a Kafka Topic, applies filters and prints the flows as JSON. The topic is either the belWue general topic or more likely a customer specific topic for one customer ID only.

## Connecting

```
-kafka.brokers string
    Kafka brokers separated by commas (default "127.0.0.1:9092,[::1]:9092")
-kafka.consumer_group string
    Kafka Consumer Group (default "dashboard")
-kafka.pass string
    Kafka password to authenticate with
-kafka.topic string
    Kafka topic to consume from (default "flow-messages-enriched")
-kafka.user string
    Kafka username to authenticate with
-kafka.auth bool
    Disable/Enable authentication to kafka (default true/enabled)
-kafka.tls bool
    Disable/Enable tls encryption to kafka (default true/enabled)
```

## Filters

```
-filter.IPsv4 string
    If defined, only flows to/from this IP V4 subnet are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple IP subnets.
-filter.IPsv6 string
    If defined, only flows to/from this IP V6 subnet are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple IP subnets.
-filter.customerid string
    If defined, only flows for this customer are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple customers.
-filter.peers string
    If defined, only flows to/from this peer are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple peers.
```

Example: `dumper [... connection options ...] --filter.customerid 10109 --filter.IPsv4 134.60.XY.0/24,134.ZA.BC.128/26 --filter.peers ECIX`

## Output

```
{"ipVersion":4,"srcIP":"172.217.21.195","dstIP":"134.60.XY.236","srcPort":443,"dstPort":39132,"proto":6,"peer":"ECIX","bytes":1472,"packets":32}
{"ipVersion":4,"srcIP":"134.60.30.XX","dstIP":"172.217.22.195","srcPort":54045,"dstPort":443,"proto":6,"peer":"ECIX","bytes":5952,"packets":64}
```

## Usage with Docker Image

Use with `docker run`

```
docker run \
    -e KAFKA_BROKERS="BELWUE_KAFKA_CLUSTER" \
    -e KAFKA_AUTH="true" \
    -e KAFKA_TLS="true" \
    -e KAFKA_TOPIC="flow-messages-enriched-YOURCID" \
    -e KAFKA_CONSUMER_GROUP="YOURCID-DUMPER" \
    -e KAFKA_USER="YOUR_USERNAME" \
    -e KAFKA_PASS="" \
    -e FILTER_CUSTOMERIDS="" \
    -e FILTER_IPSV4="134.60.0.0/16" \
    -e FILTER_IPSV6="" \
    -e FILTER_PEERS="DFN Stuttgart,DFN Karlsruhe" \
    omi-registry.e-technik.uni-ulm.de/bwnetflow/kafka/consumer_dumper:latest
```

via docker-compose

```
version: '2'
services:
  kafka:
    image: omi-registry.e-technik.uni-ulm.de/bwnetflow/kafka/consumer_dumper:latest
    environment:
        KAFKA_BROKERS: ...
        KAFKA_AUTH: false
        KAFKA_TLS: false
        KAFKA_TOPIC: enriched_goflow_topic
        KAFKA_CONSUMER_GROUP: myconsumer
        FILTER_CUSTOMERIDS: 
        FILTER_IPSV4: 
        FILTER_IPSV6: 
        FILTER_PEERS: 
        

```