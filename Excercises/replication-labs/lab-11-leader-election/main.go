package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-11-leader-election: majority vote elects leader for term 1")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("n1", "n2", "n3")

	term := "1"
	h.Net.Send(shared.Message{From: "n2", To: "n1", Body: map[string]string{"vote:term": term, "vote:for": "n2"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "n3", To: "n1", Body: map[string]string{"vote:term": term, "vote:for": "n2"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "election", To: "n2", Body: map[string]string{"leader": "n2", "term": term}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "election", To: "n1", Body: map[string]string{"leader": "n2", "term": term}}, 8*time.Millisecond)
	h.Net.Send(shared.Message{From: "election", To: "n3", Body: map[string]string{"leader": "n2", "term": term}}, 8*time.Millisecond)

	time.Sleep(35 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
