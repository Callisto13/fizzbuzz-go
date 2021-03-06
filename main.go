package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/callisto13/fizzbuzz-go/pkg/fizzbuzz"
)

func main() {
	numbersAsStrings := os.Args[1:]

	if len(numbersAsStrings) == 0 {
		fmt.Println("usage: fizzbuzz <number> [<number> <number> ...]")
		os.Exit(1)
	}

	for _, value := range numbersAsStrings {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("input is not a number, skipping")
			continue
		}

		fmt.Println(fizzbuzz.Says(number))
	}
}
