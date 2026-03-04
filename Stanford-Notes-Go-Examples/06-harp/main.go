package main

import "fmt"

func main() {
	fmt.Println("Harp: primary-backup with witness-assisted safety")
	primaryLog := []string{"op1", "op2"}
	backupLog := []string{"op1", "op2"}
	witnessView := 7
	fmt.Printf("primary=%v backup=%v witnessView=%d\n", primaryLog, backupLog, witnessView)
	fmt.Println("new primary must present up-to-date log for recovery")
}
