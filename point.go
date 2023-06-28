package main

import "fmt"

func main() {
	var price int = 999
	println(price)

	// var addr *int = &price
	addr := &price

	// fmt.Println(addr)
	*addr = 1000 // * จะเห็นได้บ่อย

	fmt.Println(price)
	// println(price)

}
