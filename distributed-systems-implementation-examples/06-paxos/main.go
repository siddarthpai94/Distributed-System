package main

import "fmt"

type acceptor struct {
	promised int
	accepted string
}

func main() {
	fmt.Println("Paxos: single-decree choose value")
	accs := []acceptor{{0, ""}, {0, ""}, {0, ""}}
	proposalN, value := 1, "x=42"
	promises := 0
	for i := range accs {
		if proposalN > accs[i].promised {
			accs[i].promised = proposalN
			promises++
		}
	}
	if promises >= 2 {
		for i := range accs {
			accs[i].accepted = value
		}
		fmt.Printf("chosen value: %s (majority accepted)\n", value)
	}
}
