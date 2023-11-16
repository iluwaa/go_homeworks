package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const validFormat = `^(\d+[+-\\*])+\d+$`
const splitRegExp = `(\d+|.)`

func Calculate(expr string) (float64, error) {

	exprValidated, err := ValidateExpression(expr)
	if err != nil {
		return 0, err
	}

	regexpSplit := regexp.MustCompile(splitRegExp)
	exprSplitted := regexpSplit.FindAllString(exprValidated, -1)

	result, err := strconv.ParseFloat(exprSplitted[0], 64)
	if err != nil {
		return 0, err
	}

	for i := 1; i < len(exprSplitted); i += 2 {
		operation := exprSplitted[i]
		number, err := strconv.ParseFloat(exprSplitted[i+1], 64)
		if err != nil {
			return 0, err
		}

		switch operation {
		case "*":
			result *= number
		case "/":
			if number == 0 {
				return 0, errors.New("Division by zero error!")
			}
			result /= number
		case "+":
			result += number
		case "-":
			result -= number
		}
	}

	return result, nil

}

func ValidateExpression(expr string) (string, error) {
	trimmedExpr := strings.Join(strings.FieldsFunc(expr, func(c rune) bool { return unicode.IsSpace(c) }), "")

	regexpValidate := regexp.MustCompile(validFormat)

	if len(regexpValidate.FindAllString(trimmedExpr, -1)) == 0 {
		return "", errors.New("Expression should start and ends with number and contain only numbers, '+', '-', '/', '*'.")
	}

	return trimmedExpr, nil
}

// unsuccessful attempt to order operations
func orderExpression(expr string) string {
	regexpSplit := regexp.MustCompile(splitRegExp)

	exprSplitted := regexpSplit.FindAllString(expr, -1)

	exprOrderedSlice := make([]string, 0)

	for index, value := range exprSplitted {
		fmt.Printf("index: %d, value: ", index)
		fmt.Println(value)
		if (value == "*" || value == "/") && len(exprOrderedSlice) == 0 {
			exprOrderedSlice = append(exprOrderedSlice, exprSplitted[index-1]+exprSplitted[index]+exprSplitted[index+1])

		} else if value == "*" || value == "/" {
			exprOrderedSlice = append(exprOrderedSlice, exprSplitted[index]+exprSplitted[index+1])
		}
	}

	for index, value := range exprSplitted {
		fmt.Printf("index: %d, value: ", index)
		fmt.Println(value)
		if (value == "+" || value == "-") && index-1 == 0 {
			exprOrderedSlice = append(exprOrderedSlice, "-"+exprSplitted[index-1]+exprSplitted[index]+exprSplitted[index+1])
			fmt.Println(exprOrderedSlice, exprSplitted[index-1]+exprSplitted[index]+exprSplitted[index+1])

		} else if value == "+" || value == "-" {
			exprOrderedSlice = append(exprOrderedSlice, exprSplitted[index]+exprSplitted[index+1])
		}
	}

	exprOrdered := strings.Join(exprOrderedSlice, "")
	return (exprOrdered)
}
