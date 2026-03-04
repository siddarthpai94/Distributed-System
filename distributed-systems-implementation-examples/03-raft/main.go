package main

import "fmt"

func main() {
	fmt.Println("Raft: leader appends entry and commits on majority")
	followersAck := 2
	log := []string{"set k=v"}
	if followersAck+1 >= 2 { // leader + followers >= majority of 3
		fmt.Printf("committed index 1: %s\n", log[0])
	} else {
		fmt.Println("not enough acknowledgements")
	}
}
