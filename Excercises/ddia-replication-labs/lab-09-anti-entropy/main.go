package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-09-anti-entropy: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"k": "orig"}}, 5*time.Millisecond)
	// simulate corruption by directly writing to one node's store (for demo)
	time.Sleep(30 * time.Millisecond)
	if n, ok := h.Nodes["n2"]; ok {
		n.Store.Put("k", "corrupt")
	}
	time.Sleep(10 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
