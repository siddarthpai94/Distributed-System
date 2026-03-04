package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Spanner: TrueTime commit-wait sketch")
	now := time.Now()
	epsilon := 5 * time.Millisecond
	commitTS := now.Add(2 * time.Millisecond)
	wait := commitTS.Add(epsilon).Sub(now)
	fmt.Printf("commitTS=%s epsilon=%s wait=%s\n", commitTS.Format(time.RFC3339Nano), epsilon, wait)
	fmt.Println("after commit-wait, external consistency is preserved")
}
