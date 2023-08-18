package main

import "fmt"

// interface คือ ข้อตกลงร่วมกันเพื่อให้ทำงานร่วมกันได้
// นิยามชื่อใน interface
//type promotion interface {
//	discount() int // ข้อตกลง
//}

type promotion interface {
	discount() int // ข้อตกลง
}

type infoer interface {
	info()
}

type presenter interface {
	promotion
	infoer
}

type course struct {
}

// course มีพฤติกรรม discount
func (c course) discount() int {
	return 599
}

func (c course) info() {
	fmt.Println("info: ", c)
}

// แล้วคราวนี้จะทำอย่างไงอยากตั้งชื่อให้กับ interface
func sale(val promotion) {
	// ถ้าเกิดเราให้ type เป็น interface นั้นๆ เราก็จะเห็นแค่ของใน interface เท่านั้น ไม่เห็นตัวอื่นเพราะไม่ได้อยู่ในข้อตกลง
	fmt.Printf("sale: %#v\n", val.discount())
	//	เราควรจะเลือก info ได้สิ แต่ไม่ได้ถ้าไม่มี info
}

func showInfo(val infoer) {
	val.info()
}

func summary(val presenter) {
	fmt.Println("discount price: ", val.discount())
	val.info()
}

// golang จะรู้ในขณะ compile time ว่ามีพฤติกรรมเป็นไปถาม interface ตรงไปตาม interface
func main() {
	v := course{}

	// ใน main
	//v.info()
	sale(v)
	showInfo(v)

	summary(v)
}
