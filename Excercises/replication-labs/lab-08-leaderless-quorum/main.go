package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-08-leaderless-quorum: n=5, w=3, r=3 with stale replica")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("n1", "n2", "n3", "n4", "n5")

	// Write reaches quorum nodes n1,n2,n3.
	write := map[string]string{"product:9": "v2", "version:product:9": "2"}
	h.Net.Send(shared.Message{From: "client", To: "n1", Body: write}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "client", To: "n2", Body: write}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "client", To: "n3", Body: write}, 5*time.Millisecond)

	// Stale node n4 keeps old version.
	h.Net.Send(shared.Message{From: "old", To: "n4", Body: map[string]string{"product:9": "v1", "version:product:9": "1"}}, 5*time.Millisecond)
	time.Sleep(25 * time.Millisecond)

	// Read quorum probes n2,n4,n5 -> would pick v2 by version.
	fmt.Println("read quorum sample values: n2,n4,n5")
	for _, id := range []string{"n2", "n4", "n5"} {
		v, _ := h.Nodes[id].Store.Get("product:9")
		ver, _ := h.Nodes[id].Store.Get("version:product:9")
		fmt.Printf("%s -> value=%q version=%q\n", id, v, ver)
	}

	shared.PrintClusterState(h.Nodes)
}
