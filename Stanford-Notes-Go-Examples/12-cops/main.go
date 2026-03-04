package main

import "fmt"

type write struct {
	key  string
	deps []string
}

func main() {
	fmt.Println("COPS: causal dependencies before visibility")
	w := write{key: "comment:9", deps: []string{"post:3@v5"}}
	fmt.Printf("write=%s dependsOn=%v\n", w.key, w.deps)
	fmt.Println("replica exposes write only after dependencies are present")
}
