package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-07-conflict-resolution: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("r1")
	h.AddNode("r2")
	h.AddNode("r3")

	// simulate concurrent writes by sending different values to two nodes
	h.Net.Send(shared.Message{From: "c1", To: "r1", Body: map[string]string{"item": "A"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c2", To: "r2", Body: map[string]string{"item": "B"}}, 5*time.Millisecond)
	time.Sleep(50 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
