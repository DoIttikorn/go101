package basic

import "fmt"

func zerovalue() {
	var i int
	var f float64
	var b bool
	var s string
	var r rune
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, r)
}
