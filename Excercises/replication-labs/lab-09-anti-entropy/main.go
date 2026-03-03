package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-09-anti-entropy: detect divergence and repair stale replica")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("n1", "n2", "n3")

	good := map[string]string{"k": "correct", "hash:k": "h-correct"}
	h.Net.Send(shared.Message{From: "client", To: "n1", Body: good}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "n1", To: "n2", Body: good}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "n1", To: "n3", Body: good}, 5*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Corruption on n2.
	h.Nodes["n2"].Store.Put("k", "corrupt")
	h.Nodes["n2"].Store.Put("hash:k", "h-corrupt")

	// Anti-entropy cycle repairs n2 from n1 authoritative data.
	h.Net.Send(shared.Message{From: "repair", To: "n2", Body: good}, 5*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
