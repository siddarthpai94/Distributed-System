package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-02-replication-log: append ordered entries and replay on followers")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader", "f1", "f2")

	wal := []map[string]string{
		{"seq:1": "PUT k1=v1", "k1": "v1"},
		{"seq:2": "PUT k2=v2", "k2": "v2"},
		{"seq:3": "PUT k3=v3", "k3": "v3"},
	}

	for _, entry := range wal {
		h.Net.Send(shared.Message{From: "client", To: "leader", Body: entry}, 5*time.Millisecond)
		h.Net.Send(shared.Message{From: "leader", To: "f1", Body: entry}, 8*time.Millisecond)
		h.Net.Send(shared.Message{From: "leader", To: "f2", Body: entry}, 8*time.Millisecond)
	}

	time.Sleep(60 * time.Millisecond)
	shared.PrintClusterState(h.Nodes)
}
