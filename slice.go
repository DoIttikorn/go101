package main

import "fmt"

func main() {
	var x []int
	fmt.Println("Default value of a slice is", x)
	fmt.Println(nil == x)
	fmt.Printf("Default value of a slice is%#v\n", x)
	fmt.Println("Length of an empty slice is", len(x))
	fmt.Println("Capacity of an empty slice is", cap(x))
}

// ในภาษา Go, nil คือค่า zero value สำหรับบางประเภทของข้อมูล โดยใช้เพื่อแสดงว่าไม่มีค่าหรือไม่มีการกำหนดค่าให้กับตัวแปรนั้นๆ ในกรณีของ slice, map, pointer, channel, interface, และ function ที่ไม่ได้กำหนดค่ามา หรือยังไม่มีการกำหนดค่า จะถือว่าเป็น nil
