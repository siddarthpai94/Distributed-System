package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-01-single-leader: leader accepts writes and replicates to followers")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader", "follower-1", "follower-2")

	// Client writes to leader.
	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{"user:1": "alice"}}, 5*time.Millisecond)

	// Leader fan-out replication (explicit in scaffold for clarity).
	h.Net.Send(shared.Message{From: "leader", To: "follower-1", Body: map[string]string{"user:1": "alice"}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "follower-2", Body: map[string]string{"user:1": "alice"}}, 8*time.Millisecond)

	time.Sleep(40 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
