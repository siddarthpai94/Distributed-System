package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Introduction: timeout + retry in unreliable RPC")
	attempts := 3
	for i := 1; i <= attempts; i++ {
		fmt.Printf("attempt %d: sending request\n", i)
		time.Sleep(30 * time.Millisecond)
		if i < 3 {
			fmt.Println("timeout; retrying")
			continue
		}
		fmt.Println("success on retry")
	}
}
