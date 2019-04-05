package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adibrastegarnia/ygot-scripts/src/pkg/topo"
	"github.com/openconfig/ygot/ygot"
)

func readTopoFile() (string, interface{}, interface{}) {
	jsonFile, err := os.Open("testdata/sample_topo.json")
	// If os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// Convert to byteArray
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var root map[string]interface{}

	json.Unmarshal([]byte(byteValue), &root)
	networksMap := root["ietf-network:networks"].(map[string]interface{})
	networkMap := networksMap["network"].([]interface{})[0].(map[string]interface{})

	// Extract the network-id
	networkID := fmt.Sprintf("%v", networkMap["network-id"])

	// Extract list of the nodes
	nodes := networkMap["node"]

	// Extract list of the links
	links := networkMap["ietf-network-topology:link"]

	return networkID, nodes, links
}

func createTopo(networkID string, nodes interface{}, links interface{}) topo.IETFNetwork_Networks {
	var networks topo.IETFNetwork_Networks

	// Create a network topology
	var network = networks.GetOrCreateNetwork(networkID)

	// Add nodes to the network topology
	nodesInterface := nodes.([]interface{})
	for _, value := range nodesInterface {
		nodeValues := value.(map[string]interface{})
		nodeID := fmt.Sprintf("%v", nodeValues["node-id"])
		network.NewNode(nodeID)

		// Add termination points to the nodes
		terminationPoints := nodeValues["termination-point"].([]interface{})
		for key := range terminationPoints {
			terminationPoint := terminationPoints[key].(map[string]interface{})
			terminationPointID := fmt.Sprintf("%v", terminationPoint["tp-id"])
			node := network.GetNode(nodeID)
			node.NewTerminationPoint(terminationPointID)
		}
	}

	// Add links to the network topology
	linksInterface := links.([]interface{})
	for _, value := range linksInterface {
		linkValues := value.(map[string]interface{})
		// Creates a link using the linkID
		linkID := fmt.Sprintf("%v", linkValues["link-id"])
		network.NewLink(linkID)

		// Add source and destination termination points to the links
		source := linkValues["source"].(map[string]interface{})
		destination := linkValues["destination"].(map[string]interface{})

		for key, value := range source {

			if key == "source-node" {
				link := network.GetLink(linkID)
				src := link.GetOrCreateSource()
				src.SourceNode = ygot.String(value.(string))

			}

			if key == "source-tp" {
				link := network.GetLink(linkID)
				src := link.GetOrCreateSource()
				src.SourceTp = ygot.String(value.(string))
			}

		}

		for key, value := range destination {

			if key == "dest-node" {
				link := network.GetLink(linkID)
				dest := link.GetOrCreateDestination()
				dest.DestNode = ygot.String(value.(string))
			}

			if key == "dest-tp" {
				link := network.GetLink(linkID)
				dest := link.GetOrCreateDestination()
				dest.DestTp = ygot.String(value.(string))

			}
		}

	}

	return networks
}

func main() {

	// Read JSON topology file
	networkID, nodes, links := readTopoFile()

	// Create topology from the list of nodes, links, and networkID
	networks := createTopo(networkID, nodes, links)

	var network = networks.GetOrCreateNetwork(networkID)

	// A sample test
	fmt.Println(network.GetLink("D1,1-2-1,D2,2-1-1").GetLinkId())

}
