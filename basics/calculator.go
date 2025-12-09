package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an expression: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "quit" {
			break
		}

		parseCommand(input)
	}
}

func calculate(a int, operator string, b int) {
	var result int

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		if b == 0 {
			fmt.Println("Error: cannot divide by 0")
			return
		}
		result = a / b
	case "*":
		result = a * b
	default:
		fmt.Println("Error: Uknown operator, please enter a valid operator")
		return
	}

	fmt.Println(result)
}

func parseCommand(input string) {
	parts := strings.Fields(input)

	if len(parts) < 3 {
		fmt.Println("Usage: <number1> <operator> <number2>")
		fmt.Println("Example: 5 + 5")
		return
	}

	numStr1 := parts[0]
	operator := parts[1]
	numStr2 := parts[2]

	num1, err1 := strconv.Atoi(numStr1)
	num2, err2 := strconv.Atoi(numStr2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please enter valid numbers")
		return
	}

	calculate(num1, operator, num2)
}
