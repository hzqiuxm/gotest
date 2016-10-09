package main

import (
	"fmt"
	"../stringutils"
)

/**
my first program in go !
 */
func main() {
	fmt.Println("horld world !")
	fmt.Println("hello go!")

	var i int;
	var f float32;
	i = 8
	f = 7
	fmt.Println("i == ", i)
	fmt.Println("f == ", f)
	fmt.Printf("i's type is %T\n", i)
	fmt.Printf("f's type is %T\n", f)

	var a, b = 1, "foo"
	fmt.Println("a == ", a, "b == ", b)
	fmt.Printf("a's type is %T\n", a)
	fmt.Printf("b's type is %T\n", b)

	var ptr *int
	ptr = &i
	fmt.Println("ptr == ", ptr)
	fmt.Println("*ptr == ", *ptr)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("type of x :%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case func(int) float64:
		fmt.Printf("x is func(int)")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("don't know the type")
	}

	fmt.Println(stringutils.Reverse("!oG ,olleH"));



}