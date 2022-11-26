package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstNumber, _ := getInput("Enter first number: ", reader)
	secondNumber, _ := getInput("Enter second number: ", reader)
	parsedFirstNumber, parsedSecondNumber := parseNumbers(firstNumber, secondNumber)

	option, _ := getInput("a (addition) - s (substraction) - d (division) - m (multiplication): ", reader)

	chooseOption(option, parsedFirstNumber, parsedSecondNumber)
}
