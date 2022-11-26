package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumbers() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	firstNumber, _ := getInput("Enter first number: ", reader)
	secondNumber, _ := getInput("Enter second number: ", reader)

	return firstNumber, secondNumber
}

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

func chooseOption(x float64, y float64) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("a (addition) - s (substraction) - d (division) - m (multiplication): ", reader)

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
