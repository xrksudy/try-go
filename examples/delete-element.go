package main

import "fmt"

func main() {
	// 2 ways to delete an element from a slice
	// 1. Fast version(changes order)
	a := []string{"A", "B", "C", "D"}
	i := 2

	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	fmt.Println(a)

	// 2. Slow version(maintains order)
	b := []string{"A", "B", "C", "D", "E"}
	j := 2

	copy(b[j:], b[j+1:])
	b[len(b)-1] = ""
	b = b[:len(b)-1]
	fmt.Println(b)

}
