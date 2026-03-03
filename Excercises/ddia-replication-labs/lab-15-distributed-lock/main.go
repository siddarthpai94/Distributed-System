package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-15-distributed-lock: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	// demo: store a lock-like key
	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"lock:resource": "owner1"}}, 5*time.Millisecond)
	time.Sleep(40 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
