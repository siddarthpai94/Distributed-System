package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-13-raft-full: leader failover and new term commit")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	leader := h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	// Term 1 leader writes.
	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"term": "1", "k": "v1"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "n1", To: "n2", Body: map[string]string{"term": "1", "k": "v1"}}, 5*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Leader fails, term 2 leader takes over.
	leader.Stop()
	h.Net.Send(shared.Message{From: "election", To: "n2", Body: map[string]string{"leader": "n2", "term": "2"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "client", To: "n2", Body: map[string]string{"term": "2", "k": "v2"}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "n2", To: "n3", Body: map[string]string{"term": "2", "k": "v2"}}, 8*time.Millisecond)
	time.Sleep(35 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
