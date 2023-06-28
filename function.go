package main

import "fmt"

// signature ของ function = func(a, b int) int
var plus2 = func(a, b int) int {
	return a + b
}

var minus2 = func(a, b int) int {
	return a - b
}

func main() {
	// result, _ := plus(1, 2, 3)

	// println(result)
	// printNumber(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// operator(minus2)

	price := 1000
	fmt.Printf("%p\n", &price)
	fmt.Println(price)
	convertA(price, 2)
	fmt.Println(price)
}

func convertA(price, rate int) {
	fmt.Printf("%p\n", &price)
	price = price * rate
}

func operator(process func(int, int) int) {
	fmt.Println(process(1, 2))

}

func plus(a, b, v int) (int, string) {
	return a + b, "Hello"
}

// นิยมใช้ตอน function สั้นๆ
func minus(a, b int) (result int) {
	result = a - b

	return
}

func printNumber(num1 int, nums ...int) {
	fmt.Println("num1", num1)
	for _, v := range nums {
		fmt.Println(v)
	}
}
