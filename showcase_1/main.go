package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Need 2 arguments")
		os.Exit(1)
	}

	firstVal, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	secondVal, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if firstVal < secondVal {
		fmt.Println("The second value cannot be larger than the first!")
		os.Exit(1)
	}

	result, remainder := calculate(firstVal, secondVal)

	fmt.Println("Result:", result, "Remainder:", remainder)
}

func calculate(val1 int, val2 int) (int, int) {
	return val1 / val2, val1 % val2
}
