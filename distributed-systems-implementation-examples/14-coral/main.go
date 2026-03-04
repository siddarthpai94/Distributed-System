package main

import "fmt"

func main() {
	fmt.Println("Coral: ownership transfer with versioned metadata")
	owner := "dc-west"
	version := 10
	fmt.Printf("initial owner=%s version=%d\n", owner, version)
	owner = "dc-east"
	version++
	fmt.Printf("after transfer owner=%s version=%d\n", owner, version)
}
