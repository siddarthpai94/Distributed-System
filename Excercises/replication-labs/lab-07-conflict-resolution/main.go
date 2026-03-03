package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-07-conflict-resolution: concurrent writes and deterministic merge policy")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader-a", "leader-b", "replica")

	h.Net.Send(shared.Message{From: "c1", To: "leader-a", Body: map[string]string{"item": "A@t1"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c2", To: "leader-b", Body: map[string]string{"item": "B@t1"}}, 5*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// LWW-like resolution chosen by policy here.
	resolved := map[string]string{"item": "B@t1", "conflict:item": "A@t1|B@t1"}
	h.Net.Send(shared.Message{From: "resolver", To: "leader-a", Body: resolved}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "resolver", To: "leader-b", Body: resolved}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "resolver", To: "replica", Body: resolved}, 5*time.Millisecond)
	time.Sleep(30 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
