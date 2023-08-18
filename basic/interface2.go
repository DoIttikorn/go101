package basic

import "fmt"

type promotion interface {
	discount() int
}

type priceProduct float32

// type course struct{}

// func (c course) discount() int {
// 	return 100000
// }

// func (c course) info() string {
// 	return "Golang"
// }

func (p priceProduct) discount() int {
	return 8888
}

func sale(v promotion) {
	fmt.Printf("sale: %#v\n", v.discount())
	// fmt.Printf("info: %#v\n", v.info())
}

func interface2() {
	v := priceProduct(1000.00)

	sale(v)
}
