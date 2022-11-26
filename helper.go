package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumbers() (float64, float64) {
	reader := bufio.NewReader(os.Stdin)

	firstNumber, _ := getInput("Enter first number: ", reader)
	secondNumber, _ := getInput("Enter second number: ", reader)

	parsedFirstNumber, err1 := strconv.ParseFloat(firstNumber, 64)
	parsedSecondNumber, err2 := strconv.ParseFloat(secondNumber, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("You entered an invalid number, please try again.")
		getNumbers()
	}

	return parsedFirstNumber, parsedSecondNumber
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func chooseOption(x float64, y float64) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput(
		`
a - (addition)
s - (substraction)
d - (division)
m - (multiplication)

Choose an option: `, reader)

	switch strings.ToLower(option) {
	case "a":
		addition(x, y)
	case "s":
		substraction(x, y)
	case "d":
		division(x, y)
	case "m":
		multiplication(x, y)
	default:
		fmt.Println("Please enter a valid option.")
		chooseOption(x, y)
	}
}
