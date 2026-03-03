package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func main() {
	fmt.Println("lab-15-distributed-lock: single owner lock with lease renewal")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("coordinator", "n2", "n3")

	// Client A acquires lock.
	h.Net.Send(shared.Message{From: "client-a", To: "coordinator", Body: map[string]string{"lock:resource": "client-a", "lease_ms": "1500"}}, 5*time.Millisecond)
	time.Sleep(15 * time.Millisecond)

	// Client B attempt is rejected (simulated by writing rejection marker).
	h.Net.Send(shared.Message{From: "coordinator", To: "n2", Body: map[string]string{"lock:resource:rejected": "client-b"}}, 5*time.Millisecond)

	// Lease renewal from client A.
	h.Net.Send(shared.Message{From: "client-a", To: "coordinator", Body: map[string]string{"lock:resource": "client-a", "lease_ms": "1500"}}, 5*time.Millisecond)
	time.Sleep(25 * time.Millisecond)

	shared.PrintClusterState(h.Nodes)
}
