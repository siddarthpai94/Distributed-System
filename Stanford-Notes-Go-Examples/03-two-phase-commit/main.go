package main

import "fmt"

type participant struct {
	name      string
	canCommit bool
}

func main() {
	fmt.Println("2PC: prepare then commit/abort")
	parts := []participant{{"A", true}, {"B", false}, {"C", true}}
	allYes := true
	for _, p := range parts {
		fmt.Printf("prepare %s -> %v\n", p.name, p.canCommit)
		allYes = allYes && p.canCommit
	}
	if allYes {
		fmt.Println("decision: COMMIT")
	} else {
		fmt.Println("decision: ABORT")
	}
}
