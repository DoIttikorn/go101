package main

import "fmt"

func main() {
	const hello string = "Hello, World!"

	// the init statement: executed before the first iteration
	// the condition expression: evaluated before every iteration
	// the post statement: executed at the end of every iteration

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// for range
	for _, v := range hello {
		fmt.Printf("%c\n", v)
	}

	// The init and post statements are optional.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// forever
	// for {
	// }
}
