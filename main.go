package main

import (
	"fmt"
	"unicode/utf8"
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
	fmt.Printf("\n %#v\n", y_lenX)
	name := "abc"
	println(" len name =", len(name))
	password := "191144zs"
	//and slide
	//index

	slide := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	println(utf8.RuneCountInString(password))
	//want slide index 1 and lested
	fmt.Printf("%#v\n", slide[1:])
	//want output slide index 0 into 7
	fmt.Printf("%#v\n", slide[0:7])
	// using map
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["en"] = "United State"

	println(countries["th"])
	//using map get well

	map_x,map_y := 10,20

	
}
