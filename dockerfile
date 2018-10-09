FROM ubuntu:latest
RUN apt-get update ; apt-get install -y ca-certificates
WORKDIR /
ADD dumper /
ADD docker-init /
RUN chmod +x /dumper
ENV KAFKA_BROKERS="127.0.0.1:9092,[::1]:9092"
ENV KAFKA_TOPIC="flow-messages-enriched"
ENV KAFKA_CONSUMER_GROUP=""
ENV KAFKA_USER=""
ENV KAFKA_PASS=""
ENV FILTER_CUSTOMERIDS=""
ENV FILTER_IPSV4=""
ENV FILTER_IPSV6=""
ENV FILTER_PEERS=""
ENTRYPOINT [ "/docker-init" ]