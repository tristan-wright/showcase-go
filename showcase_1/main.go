package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var result float32

	operatorPtr := flag.String("operator", "add", "A string to either add, subtract, multiply or divide")
	flag.Parse()

	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Need 3 arguments\n" +
			"In the format of number operand number. E.g. 4 times 5")
		os.Exit(1)
	}

	firstVal, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	secondVal, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if firstVal < secondVal {
		fmt.Println("The second value cannot be larger than the first!")
		os.Exit(1)
	}

	err = check_operator(operatorPtr)

	if err != nil {
		fmt.Println("Incorrect operator. See help.")
	}

	switch *operatorPtr {
	case "add":
		result = add(firstVal, secondVal)

	case "divide":
		result = divide(firstVal, secondVal)

	case "multiply":
		result = multiply(firstVal, secondVal)

	case "subtract":
		result = subtract(firstVal, secondVal)
	}

	fmt.Println("Result:", result)
}

func check_operator(operator *string) error {
	valid_operators := [4]string{"add", "divide", "multiply", "subtract"}

	for _, valid_operator := range valid_operators {
		if *operator == valid_operator {
			return nil
		}
	}
	return errors.New("incorrect operator")
}

func add(val1 int, val2 int) float32 {
	return float32(val1) + float32(val2)
}

func divide(val1 int, val2 int) float32 {
	return float32(val1) / float32(val2)
}

func multiply(val1 int, val2 int) float32 {
	return float32(val1) * float32(val2)
}

func subtract(val1 int, val2 int) float32 {
	return float32(val1) - float32(val2)
}
