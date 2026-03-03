package shared

import (
	"math/rand"
	"sync"
	"time"
)

// Network simulates an unreliable network for lab experiments.
// It's intentionally tiny: message send may be delayed or dropped.
type Network struct {
	mu       sync.RWMutex
	peers    map[string]chan Message
	dropRate float32 // 0.0 - 1.0
	rand     *rand.Rand
	randMu   sync.Mutex
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

// SetDropRate updates the network drop rate atomically.
func (n *Network) SetDropRate(dropRate float32) {
	if dropRate < 0 {
		dropRate = 0
	}
	if dropRate > 1 {
		dropRate = 1
	}
	n.mu.Lock()
	n.dropRate = dropRate
	n.mu.Unlock()
}

// DropRate returns the current network drop rate atomically.
func (n *Network) DropRate() float32 {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.dropRate
}

// Send sends a message; it may be delayed or dropped.
func (n *Network) Send(msg Message, latency time.Duration) {
	n.mu.RLock()
	ch, ok := n.peers[msg.To]
	dropRate := n.dropRate
	n.mu.RUnlock()
	if !ok {
		return
	}
	// drop
	n.randMu.Lock()
	shouldDrop := n.rand.Float32() < dropRate
	n.randMu.Unlock()
	if shouldDrop {
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
