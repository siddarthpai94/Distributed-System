package shared

import (
	"math/rand"
	"sync"
	"time"
)

// Network simulates an unreliable network for lab experiments.
// It's intentionally tiny: message send may be delayed or dropped.
type Network struct {
	mu       sync.Mutex
	peers    map[string]chan Message
	dropRate float32 // 0.0 - 1.0
	rand     *rand.Rand
}

// Message is a small generic payload.
type Message struct {
	From string
	To   string
	Body interface{}
}

// NewNetwork creates a simulated network.
func NewNetwork(dropRate float32) *Network {
	return &Network{
		peers:    make(map[string]chan Message),
		dropRate: dropRate,
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Register creates a channel for a node ID.
func (n *Network) Register(nodeID string) chan Message {
	n.mu.Lock()
	defer n.mu.Unlock()
	ch := make(chan Message, 128)
	n.peers[nodeID] = ch
	return ch
}

// Send sends a message; it may be delayed or dropped.
func (n *Network) Send(msg Message, latency time.Duration) {
	n.mu.Lock()
	ch, ok := n.peers[msg.To]
	n.mu.Unlock()
	if !ok {
		return
	}
	// drop
	if n.rand.Float32() < n.dropRate {
		return
	}
	// deliver after latency asynchronously
	go func() {
		time.Sleep(latency)
		select {
		case ch <- msg:
		default:
			// if receiver busy, drop message (simulates overload)
		}
	}()
}
