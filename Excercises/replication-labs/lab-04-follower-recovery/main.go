package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-04-follower-recovery: follower misses writes, then catches up")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	leader := h.AddNode("leader")
	f1 := h.AddNode("f1")
	h.AddNode("f2")

	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{"k1": "v1"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f1", Body: map[string]string{"k1": "v1"}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f2", Body: map[string]string{"k1": "v1"}}, 8*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Simulate follower crash.
	f1.Stop()
	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{"k2": "v2"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f2", Body: map[string]string{"k2": "v2"}}, 8*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Follower restarts and receives replay/catch-up.
	f1.Start()
	h.Net.Send(shared.Message{From: "leader", To: "f1", Body: map[string]string{"k1": "v1", "k2": "v2"}}, 8*time.Millisecond)
	_ = leader
	time.Sleep(30 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
