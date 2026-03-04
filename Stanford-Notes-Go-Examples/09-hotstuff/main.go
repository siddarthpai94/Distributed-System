package main

import "fmt"

func main() {
	fmt.Println("HotStuff: chained quorum certificates")
	blocks := []string{"b1", "b2", "b3"}
	qc := map[string]bool{"b1": true, "b2": true, "b3": true}
	for _, b := range blocks {
		fmt.Printf("block %s hasQC=%v\n", b, qc[b])
	}
	fmt.Println("pipeline advances with continuous QC formation")
}
