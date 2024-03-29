package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	kafka "github.com/bwNetFlow/kafkaconnector"

	"github.com/Shopify/sarama"
)

var (
	// Kafka options
	kafkaBroker        = flag.String("kafka.brokers", "127.0.0.1:9092,[::1]:9092", "Kafka brokers separated by commas")
	kafkaInTopic       = flag.String("kafka.topic", "flow-messages-enriched", "Kafka topic to consume from")
	kafkaConsumerGroup = flag.String("kafka.consumer_group", "dashboard", "Kafka Consumer Group")

	kafkaUser = flag.String("kafka.user", "", "Kafka username to authenticate with")
	kafkaPass = flag.String("kafka.pass", "", "Kafka password to authenticate with")
	kafkaAuth = flag.Bool("kafka.auth", true, "Kafka auth enable/disable")
	kafkaTLS  = flag.Bool("kafka.tls", true, "Kafka tls connection enable/disable")

	// filtering
	filterCustomerIDs = flag.String("filter.customerid", "", "If defined, only flows for this customer are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple customers.")
	filterIPsv4       = flag.String("filter.IPsv4", "", "If defined, only flows to/from this IP V4 subnet are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple IP subnets.")
	filterIPsv6       = flag.String("filter.IPsv6", "", "If defined, only flows to/from this IP V6 subnet are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple IP subnets.")
	filterPeers       = flag.String("filter.peers", "", "If defined, only flows to/from this peer are considered. Leave empty to disable filter. Provide comma separated list to filter for multiple peers.")
)

var kafkaConn = kafka.Connector{}

func main() {
	flag.Parse()

	// catch termination signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(signals)
	go func() {
		<-signals
		shutdown(0)
	}()

	// Set kafka auth
	fmt.Printf("auth: %+v\n", *kafkaAuth)
	fmt.Printf("tls: %+v\n", *kafkaTLS)
	if *kafkaAuth {
		if *kafkaUser != "" {
			kafkaConn.SetAuth(*kafkaUser, *kafkaPass)
		} else {
			kafkaConn.SetAuthAnon()
		}
	} else {
		kafkaConn.DisableAuth()
	}

	if *kafkaTLS {
		// enabled by default
	} else {
		kafkaConn.DisableTLS()
	}

	// Establish Kafka Connection
	kafkaConn.StartConsumer(*kafkaBroker, []string{*kafkaInTopic}, *kafkaConsumerGroup, sarama.OffsetNewest)
	defer kafkaConn.Close()
	runKafkaListener()
}

func shutdown(exitcode int) {
	kafkaConn.Close()
	// return exit code
	os.Exit(exitcode)
}
