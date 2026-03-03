package main
package main

import (
    "fmt"
    "time"

    "example.com/ddia-replication-labs/shared"
)

func main() {
    fmt.Println("lab-06-multi-leader: demo starting")
    h := shared.NewTestHarness(0.0)
    defer h.Shutdown()
    h.AddNode("l1")
    h.AddNode("l2")
    h.AddNode("f1")

    h.Net.Send(shared.Message{From: "client", To: "l1", Body: map[string]string{"k": "v1"}}, 20*time.Millisecond)
    time.Sleep(60 * time.Millisecond)

    for id, n := range h.Nodes {
        fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
    }
}
