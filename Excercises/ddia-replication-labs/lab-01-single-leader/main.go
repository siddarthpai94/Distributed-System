package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-01-single-leader: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("node1")
	h.AddNode("node2")
	h.AddNode("node3")

	// send a simple write to node1 via the simulated network
	h.Net.Send(shared.Message{From: "client", To: "node1", Body: map[string]string{"k1": "v1"}}, 10*time.Millisecond)
	time.Sleep(50 * time.Millisecond)

	// print stores
	for id, n := range h.Nodes {
		fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
	}
}
