package main

import (
	"fmt"
	"time"
)

func sum1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func sum2(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	time1 := time.Now()
	c := make(chan int, 2)
	go sum1(s[:len(s)/2], c)
	go sum1(s[len(s)/2:], c)
	x := <-c
    y := <-c
	fmt.Println("sum1:", x+y)
	fmt.Println("time1:", time.Since(time1))

    time2 := time.Now()
	sum2 := sum2(s)
	fmt.Println("sum2:", sum2)
	fmt.Println("time2:", time.Since(time2))
}
