package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
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