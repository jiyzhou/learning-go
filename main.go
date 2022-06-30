package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	value1 := readOperand("1", reader)
	value2 := readOperand("2", reader)
	validOperators := []string{"+", "-", "*", "/"}
	operator := readOperator(validOperators, reader)
	result := calcuate(value1, value2, operator)
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n", result)

}

func calcuate(value1, value2 float64, operator string) float64 {
	var result float64
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 != 0 {
			result = value1 / value2
		} else {
			panic("Divided by zero\n")
		}
	}
	return result
}

func readOperand(i string, reader *bufio.Reader) float64 {
	fmt.Printf("Value %s: ", i)
	input, _ := reader.ReadString('\n')
	inputFloat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panicStr := "Value " + i + " must be a number\n"
		panic(panicStr)
	}
	return inputFloat
}

func readOperator(valid []string, reader *bufio.Reader) string {
	fmt.Printf("Select an operation ( + - * / ): ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	strings.TrimSpace(input)
	if !contains(valid, input) {
		panicStr := input + " is not a valid operator\n"
		panic(panicStr)
	}
	return input
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
