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
	// b = 345
	// c = "This is string"
	// d = 2.3455
	// // // Print each variables on one line
	// // fmt.Println(a)
	// // fmt.Println(b)
	// // fmt.Println(c)
	// // fmt.Println(d)
	// // fmt.Println("Printing all variables on one line")
	// // // print all variables on one line
	// // fmt.Println(a, b, c, d)

	// fmt.Printf("%v\t - это значение. а это тип - \t%T\n", a, a)
	// fmt.Printf("%d \t %b \t %#X \n", a[1], a[10], a[5])
	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= 10; j++ {
	// 		// fmt.Println(i * j)
	// 		// fmt.Println(i, "x", j, "=", i*j)
	// 		fmt.Printf("%v \t", i*j)
	// 	}
	// 	fmt.Printf("\n\n\n")
	// }

	// fmt.Printf("%v\t%T\n", b, b)
	// fmt.Printf("%v\t\t%T\n", c, c)
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
