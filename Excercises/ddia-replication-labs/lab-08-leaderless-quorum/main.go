package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-08-leaderless-quorum: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")
	h.AddNode("n4")
	h.AddNode("n5")

	h.Net.Send(shared.Message{From: "client", To: "n3", Body: map[string]string{"qk": "qv"}}, 10*time.Millisecond)
	time.Sleep(60 * time.Millisecond)

	for id, n := range h.Nodes {
		fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
	}
}
