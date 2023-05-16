package main

import "fmt"

func main() {

	// for loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("============")
	// while loop
	var sum int = 1
	for sum <= 10 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println("============")
	for {
		fmt.Println("loop")
		break
	}
	// for range
	str := "Hello, World!"

	for index, value := range str {
		fmt.Println("index = ", index, "value = ", string(value))
	}

	arrStr := []string{"Hello", "World", "SCB", "Connect"}

	for _, value := range arrStr {
		// fmt.Println("index = ", _, "value = ", value)
		fmt.Println("value = ", value)
	}

}
