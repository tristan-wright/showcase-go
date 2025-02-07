package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./secrets.txt")
	check(err)
	fmt.Print(string(dat))

	err2 := os.WriteFile("./secrets.safe", dat, 0777)
	check(err2)
}
