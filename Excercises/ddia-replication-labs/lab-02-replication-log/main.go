package main
package main

import (
    "fmt"
    "time"

    "example.com/ddia-replication-labs/shared"
)

func main() {
    fmt.Println("lab-02-replication-log: demo starting")
    h := shared.NewTestHarness(0.0)
    defer h.Shutdown()
    h.AddNode("node1")
    h.AddNode("node2")
    h.AddNode("node3")

    h.Net.Send(shared.Message{From: "client", To: "node1", Body: map[string]string{"key": "value"}}, 5*time.Millisecond)
    time.Sleep(30 * time.Millisecond)

    for id, n := range h.Nodes {
        fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
    }
}
