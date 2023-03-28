package main

import "fmt"

func main() {
	// var a int = 10
	a := 10

	banana, apple := 10, 20

	a = 23

	fmt.Println(a)
	fmt.Println(banana)
	fmt.Println(apple)

	check := true
	s1 := "Hello, World!"
	floatNumber := 10.5

	realNumber := 10 + 5i

	fmt.Println(check, s1, floatNumber, realNumber)

	var char rune = '😭'
	fmt.Printf("%c\n", char)

}

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
// represents a Unicode code point

// float32 float64

// complex64 complex128
