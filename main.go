package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstNumber, secondNumber := getNumbers(reader)
	parsedFirstNumber, parsedSecondNumber := parseNumbers(firstNumber, secondNumber)

	chooseOption(parsedFirstNumber, parsedSecondNumber)
}
