package main

import "fmt"

func main() {
	// Arrays
	// var a [2]string
	// a[0] = "Avik"
	// a[1] = "Shingha"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// Slices
	// names := [4]string{
	// 	"Avik",
	// 	"Fahad",
	// 	"Sizan",
	// 	"Taohid",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "Pathan"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// Slice length and capacity
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)

	/// Creating a slice with make
	a := make([]int, 5, 5) // len=5 cap=5
	a = append(a, 1)
	printSlice(a)
	b := make([]int,1,2)
	b = append(b, 1)  // if we append more elements than legth then it will create a new array with double capacity
	printSlice(b)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
