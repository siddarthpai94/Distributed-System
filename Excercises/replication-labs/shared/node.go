package shared

import (
	"sync"
)

// Node represents a basic node in the cluster. Labs build on this skeleton.
type Node struct {
	ID        string
	Inbox     chan Message
	Network   *Network
	Store     *KVStore
	Shutdown  chan struct{}
	stopOnce  sync.Once
	startOnce sync.Once
	wg        sync.WaitGroup
}

// NewNode constructs a node and registers it on the network.
func NewNode(id string, net *Network) *Node {
	n := &Node{
		ID:       id,
		Network:  net,
		Store:    NewKVStore(),
		Shutdown: make(chan struct{}),
	}
	n.Inbox = net.Register(id)
	return n
}

// Start a simple message-processing loop. Labs will replace or extend this.
func (n *Node) Start() {
	n.startOnce.Do(func() {
		n.wg.Add(1)
		go func() {
			defer n.wg.Done()
			for {
				select {
				case msg := <-n.Inbox:
					n.handleMessage(msg)
				case <-n.Shutdown:
					return
				}
			}
		}()
	})
}

func (n *Node) handleMessage(msg Message) {
	// default handler: expect map[string]string for simple prototyping
	switch body := msg.Body.(type) {
	case map[string]string:
		for k, v := range body {
			n.Store.Put(k, v)
		}
	default:
		// ignore unknown
	}
}

// Stop signals the node to stop.
func (n *Node) Stop() {
	n.stopOnce.Do(func() {
		close(n.Shutdown)
	})
	n.wg.Wait()
}
