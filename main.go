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
	fmt.Print("Value 1: ")
	value1, _ := reader.ReadString('\n')
	valueFloat1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	fmt.Print("Value 2: ")
	value2, _ := reader.ReadString('\n')
	valueFloat2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	result := valueFloat1 + valueFloat2
	result = math.Round(result*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n", valueFloat1, valueFloat2, result)

}
