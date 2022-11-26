package main

func main() {
	

	firstNumber, secondNumber := getNumbers()
	parsedFirstNumber, parsedSecondNumber := parseNumbers(firstNumber, secondNumber)

	chooseOption(parsedFirstNumber, parsedSecondNumber)
}
