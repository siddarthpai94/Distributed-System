package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("ZooKeeper: leader election via ephemeral sequential znodes")
	nodes := []string{"/election/n_0003", "/election/n_0001", "/election/n_0002"}
	sort.Strings(nodes)
	leader := nodes[0]
	fmt.Printf("candidates: %v\n", nodes)
	fmt.Printf("leader (smallest znode): %s\n", leader)
	fmt.Println("if leader session expires, next smallest becomes leader")
}
