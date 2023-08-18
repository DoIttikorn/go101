package basic

import "fmt"

func condition() {

	// switch
	// if-else

	point := 10
	if point > 10 {
		fmt.Println("point is greater than 10 point = ", point)
	} else if point == 10 {
		fmt.Println("point is 10 point = ", point)
	} else {
		fmt.Println("point is less than 10 point = ", point)
	}
	fmt.Println("point", point)

	switch point {
	case 10, 20, 30:
		fmt.Println("enter case 10")
		fmt.Println("point is 10 point = ", point)
		fallthrough
	default:
		fmt.Println("point is less than 10 point = ", point)
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func whatAmI(i any) { // any
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}
