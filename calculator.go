package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Calculator")

	for {
		fmt.Print("Enter an expression (or 'quit' to exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println("Exiting...")
			break
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func evaluateExpression(expression string) (float64, error) {
	// Split the expression into operands and operator
	tokens := strings.Fields(expression)
	if len(tokens) != 3 {
		return 0, fmt.Errorf("invalid expression")
	}

	operand1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand: %v", err)
	}

	operand2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand: %v", err)
	}

	operator := tokens[1]

	// Perform the arithmetic operation
	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		result = operand1 / operand2
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}

	return result, nil
}
