package main

import "strings"

func getInput() (string){

}

func chooseOption(option string, x float64, y float64)() {
	switch strings.ToLower(option) {
	case "a":
		addition()
	case "s":
		substraction()
	case "d":
		division()
	case "m":
		multiplication()
	}
}

func addition(x float64, y float64) (float64){
	return x + y
}

func substraction(x float64, y float64) (float64){
	return x - y
}

func division(x float64, y float64) (float64){
	return x / y
}

func multiplication(x float64, y float64) (float64){
	return x * y
}