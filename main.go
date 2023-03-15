package main

import (
	"fmt"
)

func main() {
	var x = 10.45
	z := 50
	_ = z
	var y int
	_ = y

	// var value1 int
	// var value2 int
	// value2=value2 //
	point := 50
	value := 40
	point = point - value
	if point >= 50 {
		fmt.Println(point)
		fmt.Printf("%d is pass\n", point)
	} else if point < 50 {
		println("Not pass")
	}
	fmt.Printf("hello \t %#.1f", x)
	// variables
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	arr[0][0] = 10
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("\n %#v", arr)
	fmt.Printf("\n %#v", arr2)
	list := []int{1, 2, 3}
	list = append(list, 4)
	y_lenX := len(list)
	fmt.Printf("\n %#v", list)
	fmt.Printf("\n %#v", y_lenX)

}
