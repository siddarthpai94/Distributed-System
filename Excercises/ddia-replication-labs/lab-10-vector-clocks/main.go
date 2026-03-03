package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-10-vector-clocks: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("a")
	h.AddNode("b")
	h.AddNode("c")

	h.Net.Send(shared.Message{From: "c1", To: "a", Body: map[string]string{"x": "A"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c2", To: "b", Body: map[string]string{"x": "B"}}, 5*time.Millisecond)
	time.Sleep(60 * time.Millisecond)

	for id, n := range h.Nodes {
		fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
	}
}
