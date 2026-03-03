package main

import (
	"context"
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-16-chaos-testing: random drops + partition while writes continue")
	h := shared.NewTestHarness(0.3)
	defer h.Shutdown()
	h.AddNodes("n1", "n2", "n3")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h.Partition(ctx, 80*time.Millisecond)

	for i := 1; i <= 5; i++ {
		payload := map[string]string{fmt.Sprintf("k%d", i): fmt.Sprintf("v%d", i)}
		h.Net.Send(shared.Message{From: "client", To: "n1", Body: payload}, 5*time.Millisecond)
		h.Net.Send(shared.Message{From: "n1", To: "n2", Body: payload}, 10*time.Millisecond)
		h.Net.Send(shared.Message{From: "n1", To: "n3", Body: payload}, 10*time.Millisecond)
		time.Sleep(12 * time.Millisecond)
	}

	time.Sleep(130 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
