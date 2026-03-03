package main
package main

import (
    "fmt"
    "time"

    "example.com/ddia-replication-labs/shared"
)

func main() {
    fmt.Println("lab-03-async-vs-sync: demo starting")
    h := shared.NewTestHarness(0.0)
    defer h.Shutdown()
    h.AddNode("n1")
    h.AddNode("n2")
    h.AddNode("n3")

    h.Net.Send(shared.Message{From: "client", To: "n1", Body: map[string]string{"x": "1"}}, 10*time.Millisecond)
    time.Sleep(40 * time.Millisecond)

    for id, n := range h.Nodes {
        fmt.Printf("%s store: %+v\n", id, n.Store.Snapshot())
    }
}
