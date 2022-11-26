package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseNumbers(firstNumber string, secondNumber string) (float64, float64) {
	parsedFirstNumber, _ := strconv.ParseFloat(firstNumber, 64)
	parsedSecondNumber, _ := strconv.ParseFloat(secondNumber, 64)

	return parsedFirstNumber, parsedSecondNumber
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func chooseOption(option string, x float64, y float64) {

	switch strings.ToLower(option) {
	case "a":
		addition(x, y)
	case "s":
		substraction(x, y)
	case "d":
		division(x, y)
	case "m":
		multiplication(x, y)
	}
}

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
