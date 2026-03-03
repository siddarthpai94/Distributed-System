package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-05-read-consistency: eventual vs read-after-write observation")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader", "replica-a", "replica-b")

	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{"order:7": "paid"}}, 5*time.Millisecond)

	// Read from lagging replica before replication arrives (eventual anomaly window).
	time.Sleep(7 * time.Millisecond)
	if v, ok := h.Nodes["replica-a"].Store.Get("order:7"); ok {
		fmt.Printf("eventual read: %s\n", v)
	} else {
		fmt.Println("eventual read: <missing>")
	}

	// Replicate and then read-after-write from leader.
	h.Net.Send(shared.Message{From: "leader", To: "replica-a", Body: map[string]string{"order:7": "paid"}}, 20*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "replica-b", Body: map[string]string{"order:7": "paid"}}, 20*time.Millisecond)
	if v, ok := h.Nodes["leader"].Store.Get("order:7"); ok {
		fmt.Printf("read-after-write (leader): %s\n", v)
	}

	time.Sleep(35 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
