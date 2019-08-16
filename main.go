package main

import (
	"fmt"
)

var a []int
var b int
var c string
var d float32

func main() {
	a = []int{8, 2, 3, 4, 5, 6, 7, 10, 22, 33, 55}
	min := a[0]
	max := a[0]
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	fmt.Println("Maximum of ", a, " is ", max, " Min = ", min)
	fmt.Printf("Maximum of %v is %v and minimum is %v\n", a, max, min)
}
