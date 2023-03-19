package main

import (
	"fmt"
	"goBasic/customer"
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
	// map + add value to variables
	// map_x,map_y := 10,20
	// go map + if
	country, ok := countries["jp"]
	if ok {
		println(country)
	} else {
		println("no value")
	}

	// for loop

	loop_i := []int{1, 2, 3, 4, 5, 6}
	for j := 0; j < len(loop_i); {
		println(loop_i[j])

		j++
	}
	// for range loop
	loop_x := []int{10, 20, 30, 40, 50, 60}
	for i, v := range loop_x {
		println(v, i)
	}
	// i it is  index number of loop

	//function
	//using function sum
	c, St, _ := sum(10, 20)
	println(c, St)
	// in advance function

	fx := func(a, b int) int { // it is annonamus of function
		return a + b
	}

	result := fx(10, 20)
	println(result)

	//using cal function'
	cal(add)
	cal(sub)
	// if using hello function in cal function is not working
	// using baladic function
	va := []int{1, 2, 3, 4}
	s := novaladic(va)
	println(s) // it is normal function for loop add value slides

	// if valadic function

	va2 := valadic(1, 2, 3, 4, 5)
	println(va2)
	//and fmt.Print
	fmt.Printf("%v %v", 1, 2)
	// calls package custom
	println(customer.Hello())

	// pass values between variables

	frist := 10
	second := frist

	// check values

	second = 20
	println(frist, second)

	// go pointer
	var x_pointer, y_pointer int
	x_pointer = 10
	y_pointer = x_pointer
	fmt.Println(&x_pointer)
	fmt.Println(&y_pointer)
	// add pointer to variables
	var z_pointer *int
	z_pointer = &x_pointer
	fmt.Println(z_pointer)
	// display variables z_pointer
	fmt.Println(*z_pointer)
	//add new valus  to z_pointer
	*z_pointer = 20
	// and doing x_pointer new valus
	fmt.Println(*&x_pointer)
	// using result function
	results(&x_pointer)
	fmt.Println(*&x_pointer)

}
func cal(f func(int, int) int) {
	sum := f(50, 10)
	println(sum)
}

// type function is int,string,bool
func sum(a, b int) (int, string, bool) {
	return a + b, "Hello, world!", true
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func hello(name string) string {
	return "hello, world!" + name
}

func novaladic(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func valadic(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
func results(re *int) {
	a := 20
	b := 10
	*re = a + b
}

func results2(re2 *int) {

}
