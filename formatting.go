package main

import "fmt"

func main() {
	var A rune = 'A'

	// var left rune = '◀'

	var msg, age, price, check = "Hello World", 20, 100.523, true

	fmt.Printf("%c\n", A)

	fmt.Printf("%s\n", msg)
	fmt.Printf("%d\n", age)
	fmt.Printf("%v\n", price)
	fmt.Printf("%t\n", check)

	fmt.Println("working", age, price, check)

	// need
	fmt.Printf("%T %v\n", msg, msg)
	fmt.Printf("%#v\n", msg)

	// var arrayNumber = [3]int{1, 2, 3}

	// fmt.Println(arrayNumber)
	// fmt.Printf("%T %#v\n", arrayNumber, arrayNumber)
}
