package main

import "fmt"

func main() {
	fmt.Println("Aurora: decoupled compute and replicated storage log")
	segments := []string{"seg-101", "seg-102", "seg-103"}
	quorumWritten := 4
	fmt.Printf("segments=%v replicasAck=%d\n", segments, quorumWritten)
	fmt.Println("database instance replays storage log on failover")
}
