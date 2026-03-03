package main
package main

import (
    "fmt"
    "time"

    "example.com/ddia-replication-labs/shared"
)

func main() {
    fmt.Println("lab-05-read-consistency: demo starting")
    h := shared.NewTestHarness(0.0)
    defer h.Shutdown()
    h.AddNode("a")
    h.AddNode("b")
    h.AddNode("c")

    h.Net.Send(shared.Message{From: "client", To: "a", Body: map[string]string{"k": "v"}}, 5*time.Millisecond)
    time.Sleep(30 * time.Millisecond)

    for id, n := range h.Nodes {
        fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
    }
}
