package main

import "fmt"

func main() {
	var x []int
	fmt.Println("Default value of a slice is", x)
	fmt.Println(nil == x)
	fmt.Printf("Default value of a slice is%#v\n", x)
	fmt.Println("Length of an empty slice is", len(x))
	fmt.Println("Capacity of an empty slice is", cap(x))

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:5]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Println("======show slice depend on array======")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("======append slice======")
	xx := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(xx)
	xx = append(xx, 1, 2, 3)
	fmt.Println(xx)

	fmt.Println("======copy slice======")
	dst := make([]int, len(xx))
	fmt.Println(dst)
	copy(dst, xx)
	fmt.Println(dst)
}

// ในภาษา Go, nil คือค่า zero value สำหรับบางประเภทของข้อมูล โดยใช้เพื่อแสดงว่าไม่มีค่าหรือไม่มีการกำหนดค่าให้กับตัวแปรนั้นๆ ในกรณีของ slice, map, pointer, channel, interface, และ function ที่ไม่ได้กำหนดค่ามา หรือยังไม่มีการกำหนดค่า จะถือว่าเป็น nil
