package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/callisto13/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numberAsString := os.Args[1]

	number, err := strconv.Atoi(numberAsString)
	if err != nil {
		fmt.Println("arguments must be numbers")
		os.Exit(1)
	}

	fmt.Println(fizzbuzz.Says(number))
}
