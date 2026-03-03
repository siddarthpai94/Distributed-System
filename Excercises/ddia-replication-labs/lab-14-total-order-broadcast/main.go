package main

import (
	"fmt"
	"time"

	"example.com/ddia-replication-labs/shared"
)

func main() {
	fmt.Println("lab-14-total-order-broadcast: demo starting")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNode("n1")
	h.AddNode("n2")
	h.AddNode("n3")

	h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"m": "msg1"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "client", To: "n2", Body: map[string]string{"m": "msg2"}}, 5*time.Millisecond)
	time.Sleep(60 * time.Millisecond)

	for id, n := range h.Nodes {
		fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
	}
}
