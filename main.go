package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstNumber, secondNumber := getNumbers(reader)
	parsedFirstNumber, parsedSecondNumber := parseNumbers(firstNumber, secondNumber)

	option, _ := getInput("a (addition) - s (substraction) - d (division) - m (multiplication): ", reader)

	chooseOption(option, parsedFirstNumber, parsedSecondNumber)
}
