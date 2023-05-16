package main

import "fmt"

func main() {
	var a [3]int
	fmt.Printf("original a: %#v\n", a)

	a[0] = 1
	a[1] = 2
	fmt.Printf("length: %d , capacity: %d\n", len(a), cap(a))
	fmt.Printf("original a: %#v\n", a)

}
