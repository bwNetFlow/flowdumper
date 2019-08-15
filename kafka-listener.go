package main

import (
	"log"
	"runtime/debug"

	flow "github.com/bwNetFlow/protobuf/go"
)

func runKafkaListener() {
	// initialize filters: prepare filter arrays
	initFilters()

	// handle kafka flow messages in foreground
	for {
		flow := <-kafkaConn.ConsumerChannel()
		if filterApplies(flow) {
			handleFlow(flow)
		}
	}
}

func handleFlow(flow *flow.FlowMessage) {
	// handle panic while flow processing
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered panic in handleFlow", r)
			debug.PrintStack()
			log.Printf("failed flow: %+v\n", flow)
		}
	}()

	// the only action here: dump the flow
	dumpFlow(flow)
}
