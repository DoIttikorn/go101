package basic

import "fmt"

func show(val interface{}) {

	// i, oki := val.(int) // type assertion

	// if oki {
	// 	fmt.Printf("%T : %#v\n", i, i)
	// }

	// j, okj := val.(string)

	// if okj {
	// 	fmt.Printf("%T : %#v\n", j, j)
	// }

	switch v := val.(type) {
	case int:
		fmt.Printf("%T : %#v\n", v, v+1000)
	case string:
		fmt.Printf("%T : %#v\n", v, v+" Test")
	default:
		fmt.Println("Unknown Type")
	}

}

func interface1() {
	var v any

	v = 10

	show(v)

	v = "Golang"

	show(v)

	v = 3.14
	show(v)
}
