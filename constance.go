package main

import "fmt"

type rank int

// iota คือ ตัวแปรที่เริ่มต้นที่ 0 และจะเพิ่มขึ้นทีละ 1 ทุกครั้งที่เรียกใช้ ใช้ในการสร้างค่าคงที่แบบเรียงต่อกัน
const (
	Terng rank = 1 + iota
	_          // 2
	Aung
	Dhas
	Thon
)

// advanced iota usage
const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4
	admin  = read | write | remove
)

const (
	_  = 1 << (iota * 10) // ignore the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

func main() {
	const pi float64 = 3.14

	fmt.Println(pi)

	fmt.Println(Terng)
	fmt.Println(Aung)
	fmt.Println(Dhas)
	fmt.Println(Thon)

	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)
}
