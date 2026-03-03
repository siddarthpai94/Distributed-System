package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-12-raft-log-replication: append entries and commit after majority")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader", "f1", "f2")

	entry := map[string]string{"log:idx:1": "set x=10", "term:idx:1": "4"}
	h.Net.Send(shared.Message{From: "leader", To: "f1", Body: entry}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f2", Body: entry}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "leader", Body: entry}, 1*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Commit once majority has replicated.
	commit := map[string]string{"commitIndex": "1", "x": "10"}
	h.Net.Send(shared.Message{From: "leader", To: "leader", Body: commit}, 3*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f1", Body: commit}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f2", Body: commit}, 5*time.Millisecond)
	time.Sleep(25 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
