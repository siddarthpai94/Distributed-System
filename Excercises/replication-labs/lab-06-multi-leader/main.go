package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-06-multi-leader: two leaders accept local writes and cross-replicate")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader-us", "leader-eu", "f-us", "f-eu")

	h.Net.Send(shared.Message{From: "client-us", To: "leader-us", Body: map[string]string{"region:us": "ok"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "client-eu", To: "leader-eu", Body: map[string]string{"region:eu": "ok"}}, 5*time.Millisecond)

	// Local replication.
	h.Net.Send(shared.Message{From: "leader-us", To: "f-us", Body: map[string]string{"region:us": "ok"}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader-eu", To: "f-eu", Body: map[string]string{"region:eu": "ok"}}, 8*time.Millisecond)

	// Cross-datacenter replication (higher latency).
	h.Net.Send(shared.Message{From: "leader-us", To: "leader-eu", Body: map[string]string{"region:us": "ok"}}, 40*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader-eu", To: "leader-us", Body: map[string]string{"region:eu": "ok"}}, 40*time.Millisecond)
	time.Sleep(70 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
