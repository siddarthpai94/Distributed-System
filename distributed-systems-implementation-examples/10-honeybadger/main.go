package main

import "fmt"

func main() {
	fmt.Println("HoneyBadger: asynchronous batch agreement")
	proposals := map[string][]string{
		"n1": {"tx1", "tx2"},
		"n2": {"tx3"},
		"n3": {"tx4", "tx5"},
	}
	batch := []string{}
	for _, txs := range proposals {
		batch = append(batch, txs...)
	}
	fmt.Printf("agreed batch: %v\n", batch)
}
