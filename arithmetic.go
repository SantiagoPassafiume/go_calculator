package main

import "fmt"

func addition(x float64, y float64) {
	fmt.Printf("%v + %v = %v", x, y, x+y)
}

func substraction(x float64, y float64) {
	fmt.Printf("%v - %v = %v", x, y, x-y)
}

func division(x float64, y float64) {
	fmt.Printf("%v / %v = %v", x, y, x/y)
}

func multiplication(x float64, y float64) {
	fmt.Printf("%v * %v = %v", x, y, x*y)
}