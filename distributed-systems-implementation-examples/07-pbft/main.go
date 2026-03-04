package main

import "fmt"

func main() {
	fmt.Println("PBFT: pre-prepare, prepare, commit quorum")
	n := 4
	f := 1
	prepareVotes := 3
	commitVotes := 3
	fmt.Printf("n=%d f=%d prepares=%d commits=%d\n", n, f, prepareVotes, commitVotes)
	if prepareVotes >= 2*f+1 && commitVotes >= 2*f+1 {
		fmt.Println("request executed (Byzantine quorum reached)")
	}
}
