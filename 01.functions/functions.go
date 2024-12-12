package main

import "fmt"

func add(a,b,c int) int { // return single value
	return a+b+c
}

func swap(a,b int) (int,int) { // return multiple values
	return b,a
}

func slipt(sum int) (x,y int) { // return named values, it is called naked return
	x = sum*5/9
	y = sum - x
	return
}

func main() {
	println(add(1,2,3))
	fmt.Println(swap(1,2))
	fmt.Println(slipt(18))
}