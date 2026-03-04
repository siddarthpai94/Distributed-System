package main

import "fmt"

func main() {
	fmt.Println("GFS: master tracks chunk metadata and lease holder")
	chunk := "chunk-17"
	primary := "chunkserver-2"
	replicas := []string{"chunkserver-2", "chunkserver-4", "chunkserver-8"}
	fmt.Printf("%s primary=%s replicas=%v\n", chunk, primary, replicas)
}
