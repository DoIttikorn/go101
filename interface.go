package main

import "fmt"

//func show(val int) {
//	fmt.Println(val)
//}

func show(val any) {
	// เราต้องทำ type assertion
	// https://go.dev/tour/methods/15
	//i, ok := val.(int) // แกะของออกมาดู ถ้าไม่ถูกจะเกิด runtimeError
	//
	//if ok {
	//	i = i + 2
	//	fmt.Println(i)
	//}
	//
	//j, k := val.(string) // แกะของออกมาดู ถ้าไม่ถูกจะเกิด runtimeError
	//
	//if k {
	//	j = j + "hello"
	//	fmt.Println(j)
	//}

	switch v := val.(type) {
	case int:
		i := v + 2
		fmt.Println(i)
	case string:
		j := v + "hello"
		fmt.Println(j)
	default:
		fmt.Println("not handle type.")
	}

}

func main() {
	//1 การใช้ type interface
	//var b any // empty interface คือ ไม่ระบุ type ใดๆ ไว้ | เก็บตัวแปลชนิดใดก็ได้ แต่จริงๆแล้วมันแค่ห่อของไว้อย่างในเฉยๆ
	//b = "test"
	//fmt.Printf("%T %#v/n", b, b)
	//
	//b = 37
	//fmt.Printf("%T %#v/n", b, b)

	//2 การ convert type interface
	// convert v type any to int
	//var v any
	//
	//v = 36
	//
	//show(v.(int))

	//3 เปลี่ยน type armament
	var v any

	v = "world"
	show(v)

}
