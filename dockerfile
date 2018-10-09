FROM alpine:latest
RUN apk update
RUN apk add bash ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
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