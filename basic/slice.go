package basic

import "fmt"

func slice() {
	var y [5]int
	fmt.Printf("%#v \n", y)
	var x []int

	x = append(x, 10, 20, 30, 40, 50, 60, 70, 80, 90)

	num := x[3:]
	fmt.Println(num)

	fmt.Printf("Default value of a slice is%#v\n", x)
	fmt.Println("Length of an empty slice is", len(x))
	fmt.Println("Capacity of an empty slice is", cap(x))

	str := [5]string{"Hello", "World", "SCB", "Connect"}
	fmt.Println()
	fmt.Println(str)

	aRange := str[:2]
	fmt.Println(aRange)

	bRange := str[2:]
	fmt.Println(bRange)

	bRange[0] = "SCB Connect LINE"
	fmt.Println(bRange)

	fmt.Println("=============")
	fmt.Println(str)

	// nil is default in  interface , slice , map , function , pointer , channel
}
