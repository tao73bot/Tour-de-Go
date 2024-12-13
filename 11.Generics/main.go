package main

import (
	"fmt"
)

func AddInt(a int, b int) int {
	return a + b
}

func AddFloat(a float64, b float64) float64 {
	return a + b
}

// Here, we have two functions AddInt and AddFloat that do the same thing, but one works with integers and the other with floats.
// We can use generics to create a single function that works with both types.

func Add[T int | float64](a T, b T) T {
	return a + b
}

// mapping using generics

func MapValues(values []int, mapFunc func(int) int) []int { // normal function that takes int type
	var newValues []int
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func MapValuesGen[T any](values []T, mapFunc func(T) T) []T { // generic function that takes any type
	var newValues []T
	for _, value := range values {
		newValue := mapFunc(value)
		newValues = append(newValues, newValue)
	}
	return newValues
}

// struct with generics

type User[T any] struct {
	ID int
	Name string
	Data T
}

// maps with generics

type CustomMap[T comparable, V int | string] map[T]V  // comparable is a constraint that tells the compiler that the type T must be comparable.

func main() {
	//** Additions example **//
	// resultInt := AddInt(1, 2)
	// resultFloat := AddFloat(1.1, 2.2)
	// fmt.Println(resultInt)
	// fmt.Println(resultFloat)
	// result1 := Add(1, 2)
	// fmt.Println(result1)
	// result2 := Add(1.1, 2.2)
	// fmt.Println(result2)

	//** MapValues example **//
	// values := MapValues([]int{1, 2, 3, 4, 5}, func(value int) int {
	// 	return value * 3
	// })
	// fmt.Println(values)
	// valuesGen1 := MapValuesGen([]int{1, 2, 3, 4, 5}, func(value int) int {
	// 	return value * 3
	// })
	// valuesGen2 := MapValuesGen([]string{"a", "b", "c", "d", "e"}, func(value string) string {
	// 	return value + value
	// })
	// valuesGen3 := MapValuesGen([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, func(value float64) float64 {
	// 	return value * 3
	// })
	// fmt.Println(valuesGen1)
	// fmt.Println(valuesGen2)
	// fmt.Println(valuesGen3)

	//** struct with generics example **//
	// user1 := User[int]{ID: 1, Name: "John", Data: 123}
	// user2 := User[string]{ID: 2, Name: "Jane", Data: "Hello"}
	// fmt.Println(user1)
	// fmt.Println(user2)

	//** maps with generics example **//
	m := make(CustomMap[int,string])
	m[1] = "One"
	m[2] = "Two"
	fmt.Println(m)
}
