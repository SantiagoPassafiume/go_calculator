package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstNumber, _ := getInput("Enter first number: ", reader)
	secondNumber, _ := getInput("Enter second number: ", reader)
	parsedFirstNumber, _ := strconv.ParseFloat(firstNumber, 64)
	parsedSecondNumber, _ := strconv.ParseFloat(secondNumber, 64)

	option, _ := getInput("a (addition) - s (substraction) - d (division) - m (multiplication): ", reader)

	chooseOption(option, parsedFirstNumber, parsedSecondNumber)
}