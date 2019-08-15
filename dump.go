package main

import (
	"encoding/json"
	"fmt"
	"net"

	flow "github.com/bwNetFlow/protobuf/go"
	"github.com/bwNetFlow/protobuf_helpers/go"
)

// JSONFlowMessage describes a JSON representation of a single flow
type JSONFlowMessage struct {
	IPVersion string `json:"ipVersion"`
	SrcIP     string `json:"srcIP,omitempty"`
	DstIP     string `json:"dstIP,omitempty"`
	SrcPort   uint32 `json:"srcPort"`
	DstPort   uint32 `json:"dstPort"`
	Proto     uint32 `json:"proto"`
	Peer      string `json:"peer"`
	Bytes     uint64 `json:"bytes"`
	Packets   uint64 `json:"packets"`
}

func dumpFlow(flow *flow.FlowMessage) {

	flowh := protobuf_helpers.NewFlowHelper(flow)

	// translate message to JSONFlowMessage
	jsonMsg := JSONFlowMessage{
		IPVersion: flowh.IPVersionString(),
		SrcIP:     fmt.Sprintf("%v", net.IP(flow.GetSrcAddr())),
		DstIP:     fmt.Sprintf("%v", net.IP(flow.GetDstAddr())),
		SrcPort:   flow.GetSrcPort(),
		DstPort:   flow.GetDstPort(),
		Proto:     flow.GetProto(),
		Peer:      flowh.Peer(),
		Bytes:     flow.GetBytes(),
		Packets:   flow.GetPackets(),
	}

	// produce and print JSON string
	jsonstr, err := json.Marshal(jsonMsg)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", jsonstr)
}
