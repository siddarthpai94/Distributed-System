package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-12-raft-log-replication: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"k": "v"}}, 5*time.Millisecond)
	time.Sleep(50 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
