package shared
package shared

import (
    "context"
    "log"
    "time"
)

// TestHarness provides helpers to start a set of nodes and inject faults.
type TestHarness struct {
    Nodes map[string]*Node
    Net   *Network
}

// NewTestHarness creates harness with given dropRate for network.
func NewTestHarness(dropRate float32) *TestHarness {
    net := NewNetwork(dropRate)
    return &TestHarness{Nodes: make(map[string]*Node), Net: net}
}

// AddNode creates and starts a node with the id.
func (h *TestHarness) AddNode(id string) *Node {
    n := NewNode(id, h.Net)
    h.Nodes[id] = n
    n.Start()
    return n
}

// Shutdown stops all nodes.
func (h *TestHarness) Shutdown() {
    for _, n := range h.Nodes {
        n.Stop()
    }
}

// Partition simulates isolating a set of node IDs by dropping messages to/from them.
// For simplicity, this routine increases dropRate temporarily for the harness network.
func (h *TestHarness) Partition(ctx context.Context, duration time.Duration) {
    log.Printf("simulated partition: increasing drop rate")
    old := h.Net.dropRate
    h.Net.dropRate = 1.0
    go func() {
        select {
        case <-time.After(duration):
            h.Net.dropRate = old
            log.Printf("partition healed")
        case <-ctx.Done():
            h.Net.dropRate = old
        }
    }()
}
