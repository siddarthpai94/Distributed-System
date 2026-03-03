package main

import (
	"fmt"
	"time"

	"example.com/replication-labs/shared"
)

func writeSync(h *shared.TestHarness, key, value string) time.Duration {
	start := time.Now()
	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{key: value}}, 5*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f1", Body: map[string]string{key: value}}, 20*time.Millisecond)
	h.Net.Send(shared.Message{From: "leader", To: "f2", Body: map[string]string{key: value}}, 20*time.Millisecond)
	time.Sleep(25 * time.Millisecond)
	return time.Since(start)
}

func writeAsync(h *shared.TestHarness, key, value string) time.Duration {
	start := time.Now()
	h.Net.Send(shared.Message{From: "client", To: "leader", Body: map[string]string{key: value}}, 5*time.Millisecond)
	go h.Net.Send(shared.Message{From: "leader", To: "f1", Body: map[string]string{key: value}}, 50*time.Millisecond)
	go h.Net.Send(shared.Message{From: "leader", To: "f2", Body: map[string]string{key: value}}, 50*time.Millisecond)
	return time.Since(start)
}

func main() {
	fmt.Println("lab-03-async-vs-sync: compare write latency modes")
	h := shared.NewTestHarness(0.0)
	defer h.Shutdown()
	h.AddNodes("leader", "f1", "f2")

	syncLatency := writeSync(h, "k-sync", "v1")
	asyncLatency := writeAsync(h, "k-async", "v2")
	time.Sleep(90 * time.Millisecond)

	fmt.Printf("sync write latency:  %v\n", syncLatency)
	fmt.Printf("async write latency: %v\n", asyncLatency)
	shared.PrintClusterState(h.Nodes)
}
