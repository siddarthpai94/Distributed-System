package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-10-vector-clocks: concurrent versions then conflict resolution write")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("a", "b", "c")

	// Concurrent writes with explicit vector-clock metadata.
	h.Net.Send(shared.Message{From: "c1", To: "a", Body: map[string]string{"x": "A", "vc:x": "a:1"}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c2", To: "b", Body: map[string]string{"x": "B", "vc:x": "b:1"}}, 5*time.Millisecond)
	time.Sleep(20 * time.Millisecond)

	// Client resolves siblings into C that descends from both.
	resolved := map[string]string{"x": "C", "vc:x": "a:1,b:1,c:1", "siblings:x": "A|B"}
	h.Net.Send(shared.Message{From: "c3", To: "c", Body: resolved}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c", To: "a", Body: resolved}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "c", To: "b", Body: resolved}, 5*time.Millisecond)
	time.Sleep(25 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
