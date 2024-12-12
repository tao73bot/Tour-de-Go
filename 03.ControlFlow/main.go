package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:=math.Pow(x, n); v < lim { // v is only available in if and else block
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // %g is used to print float64
	}
	return lim
}

func main() {
	sum := 0

	for i:=0; i<10; i++ {
		sum +=i
	}
	fmt.Println(sum)

	for ; sum < 1000; { // 1st and 3rd part of for loop are optional
		sum += sum
	}
	fmt.Println(sum)

	// we can use for loop as while loop
	for sum < 2000 {
		sum += sum
	}
	fmt.Println(sum)

	// Conditional
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

// Exercise: Loops and Functions

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func Sqrt(x float64) float64 {
// 	var it int = 0
// 	var z float64 = 1.0
// 	for ; it < 10 && math.Abs(z*z-math.Sqrt(x)) > 1e-10; it++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)
// 	}
// 	return z
// }
// func main() {
// 	fmt.Println(Sqrt(2))
// }


// Switch
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("When's Friday?")
// 	today := time.Now().Weekday()
// 	switch time.Friday {
// 	case today + 0:
// 		fmt.Println("Today.")
// 	case today + 1:
// 		fmt.Println("Tomorrow.")
// 	case today + 2:
// 		fmt.Println("In two days.")
// 	default:
// 		fmt.Println("Too far away.")
// 	}
// }
