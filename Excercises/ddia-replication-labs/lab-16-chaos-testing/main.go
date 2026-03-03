package main

import (
	"context"
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-16-chaos-testing: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	// simple chaos demo: partition the network briefly
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		h.Partition(ctx, 50*time.Millisecond)
	}()

	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"k": "v"}}, 5*time.Millisecond)
	time.Sleep(120 * time.Millisecond)

	cancel()

	for id, n := range h.Nodes {
		fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
	}
}
