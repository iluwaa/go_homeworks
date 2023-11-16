package main

import (
	"fmt"
	"hw13/pkg/calculator"
	"os"
)

func main() {

	// valid expression
	expression := "11+2-3*44/ 5"
	result, err := calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	// division by zero
	expression = "11+2-3*44/ 0"
	result, err = calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// invalid expression 1
	expression = "11+2-3*44/"
	result, err = calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// invalid expression 2
	expression = "11+2-3*44//5"
	result, err = calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

}
