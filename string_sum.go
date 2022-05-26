package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

//goland:noinspection GoUnusedExportedFunction
func StringSum(input string) (output string, err error) {
	if input == "" || len(input) < 1 {
		return "", fmt.Errorf("error: %w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")

	inputSlice := splitBy(input, "-+")
	if len(inputSlice) != 2 {
		return "", fmt.Errorf("error: %w", errorNotTwoOperands)
	}

	xN, yN := inputSlice[0], inputSlice[1]
	xI, yI := strings.Index(input, xN), strings.Index(input, yN)
	if xI > 0 {
		xN = string(input[xI-1]) + xN
	}
	if yI > 0 {
		yN = string(input[yI-1]) + yN
	}

	x, err := strconv.Atoi(xN)
	if err != nil {
		return "", err
	}

	y, err := strconv.Atoi(yN)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(x + y), nil
}

func splitBy(s, separates string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(separates, r)
	}
	return strings.FieldsFunc(s, splitter)
}
