package main

import (
	"encoding/json"
	"fmt"
	"net"

	flow "omi-gitlab.e-technik.uni-ulm.de/bwnetflow/bwnetflow_api/go"
)

// JSONFlowMessage describes a JSON representation of a single flow
type JSONFlowMessage struct {
	IPVersion int32  `json:"ipVersion"`
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

	// translate message to JSONFlowMessage
	jsonMsg := JSONFlowMessage{
		IPVersion: int32(flow.GetIPversion()),
		SrcIP:     net.IP(flow.GetSrcIP()).String(),
		DstIP:     net.IP(flow.GetDstIP()).String(),
		SrcPort:   flow.GetSrcPort(),
		DstPort:   flow.GetDstPort(),
		Proto:     flow.GetProto(),
		Peer:      flow.GetPeer(),
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
