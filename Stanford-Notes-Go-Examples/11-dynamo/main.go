package main

import "fmt"

func main() {
	fmt.Println("Dynamo: quorum read/write and version vectors")
	n, w, r := 3, 2, 2
	fmt.Printf("n=%d w=%d r=%d (w+r>n => overlap)\n", n, w, r)
	vector := map[string]int{"n1": 2, "n2": 1}
	fmt.Printf("version vector: %v\n", vector)
	fmt.Println("concurrent versions become siblings and require reconciliation")
}
