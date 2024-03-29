package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"

	"github.com/bwNetFlow/ip_prefix_trie"
	flow "github.com/bwNetFlow/protobuf/go"
	protobuf_helpers "github.com/bwNetFlow/protobuf_helpers/go"
)

var validCustomerIDs []int
var validPeers []string

// We have to use separate tries for IPv4 and IPv6
var validIPTrieV4, validIPTrieV6 ip_prefix_trie.TrieNode
var ipFilterSet bool

func initFilters() {
	// customer ID
	if *filterCustomerIDs != "" {
		stringIDs := strings.Split(*filterCustomerIDs, ",")
		for _, stringID := range stringIDs {
			customerID, err := strconv.Atoi(stringID)
			if err != nil {
				continue
			}
			validCustomerIDs = append(validCustomerIDs, customerID)
		}
		sort.Ints(validCustomerIDs)

		outputStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(validCustomerIDs)), ","), "[]")
		log.Printf("Filter flows for customer ids %s\n", outputStr)
	} else {
		log.Printf("No customer filter enabled.\n")
	}

	// IPs v4
	if *filterIPsv4 != "" {
		ipFilterSet = true
		stringIDs := strings.Split(*filterIPsv4, ",")
		validIPTrieV4.Insert(true, stringIDs)
		// validIPTrieV4.Print("", true)
		log.Printf("Filter flows for IPs v4: %s\n", *filterIPsv4)
	} else {
		log.Printf("No IP v4 filter enabled.\n")
	}

	// IPs v6
	if *filterIPsv6 != "" {
		ipFilterSet = true
		stringIDs := strings.Split(*filterIPsv6, ",")
		validIPTrieV6.Insert(true, stringIDs)
		log.Printf("Filter flows for IPs v6: %s\n", *filterIPsv6)
	} else {
		log.Printf("No IP v6 filter enabled.\n")
	}

	// peers
	if *filterPeers != "" {
		stringIDs := strings.Split(*filterPeers, ",")
		for _, stringID := range stringIDs {
			peer := strings.Trim(stringID, " ")
			validPeers = append(validPeers, peer)
		}
		sort.Strings(validPeers)

		// outputStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(validPeers)), ","), "[]")
		log.Printf("Filter flows for peers %v\n", validPeers)
	} else {
		log.Printf("No peer filter enabled.\n")
	}
}

func filterApplies(flow *flow.FlowMessage) bool {
	// customerID filter
	if len(validCustomerIDs) == 0 || isValidCustomerID(int(flow.GetCid())) {
		// IP subnet filter
		if !ipFilterSet || isValidIP(flow.GetSrcAddr()) || isValidIP(flow.GetDstAddr()) {
			// peer filter
			flowh := protobuf_helpers.NewFlowHelper(flow)
			if len(validPeers) == 0 || isValidPeer(flowh.Peer()) {
				return true
			}
		}
	}
	return false
}

func isValidCustomerID(cid int) bool {
	pos := sort.SearchInts(validCustomerIDs, cid)
	if pos == len(validCustomerIDs) {
		return false
	}
	return validCustomerIDs[pos] == cid
}

func isValidIP(IP []byte) bool {
	// TODO improve the following workarounds ...
	test := net.IP(IP)
	var prefix string
	if test.To4() == nil {
		// ipv6
		prefix = "64"
	} else {
		// ipv4
		prefix = "32"
	}
	ipAddrStr := net.IP(IP).String() + "/" + prefix
	ipAddr, _, _ := net.ParseCIDR(ipAddrStr)
	foundIP := false
	if ipAddr.To4() == nil {
		foundIP, _ = validIPTrieV6.Lookup(ipAddr).(bool)
	} else {
		foundIP, _ = validIPTrieV4.Lookup(ipAddr).(bool)
	}
	return foundIP
}

func isValidPeer(peer string) bool {
	pos := sort.SearchStrings(validPeers, peer)
	if pos == len(validPeers) {
		return false
	}
	return validPeers[pos] == peer
}
