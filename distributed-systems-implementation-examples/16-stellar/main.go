package main

import "fmt"

func main() {
	fmt.Println("Stellar: federated voting with quorum slices")
	slices := map[string][]string{
		"A": {"A", "B"},
		"B": {"B", "C"},
		"C": {"C", "A"},
	}
	fmt.Printf("quorum slices: %v\n", slices)
	fmt.Println("agreement emerges if slices intersect adequately")
}
