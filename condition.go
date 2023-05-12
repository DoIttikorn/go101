package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

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

	checkOs()

	whatToSaturday()

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func whatToSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func checkOs() {
	if os := runtime.GOOS; os == "darwin" {
		fmt.Println("OS X.")
	} else if os == "linux" {
		fmt.Println("Linux.")
	} else {
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
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
