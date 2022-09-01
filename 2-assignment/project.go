package main

import "fmt"

type tenNumbers []int

// not a receiver function because you are not receive anything
func newArrayOneToTen() tenNumbers {
	myNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	return myNumbers
}

// receiver function because you receive the array
func (t tenNumbers) print() {
	for _, value := range t {
		if value%2 == 0 {
			fmt.Println(value, "even")
		} else {
			fmt.Println(value, "odd")
		}
	}
}