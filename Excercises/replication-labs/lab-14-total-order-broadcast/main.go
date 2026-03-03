package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-14-total-order-broadcast: all nodes apply messages in same order")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("n1", "n2", "n3")

	ordered := []map[string]string{
		{"seq": "1", "event": "create-order"},
		{"seq": "2", "event": "charge-card"},
		{"seq": "3", "event": "send-email"},
	}

	for _, msg := range ordered {
		for _, id := range []string{"n1", "n2", "n3"} {
			h.Net.Send(shared.Message{From: "broker", To: id, Body: msg}, 5*time.Millisecond)
		}
	}

	time.Sleep(35 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
