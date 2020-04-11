package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/callisto13/fizzbuzz/pkg/fizzbuzz"
)

func main() {
	numbersAsStrings := os.Args[1:]

	for _, value := range numbersAsStrings {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("input is not a number, skipping")
			continue
		}

		fmt.Println(fizzbuzz.Says(number))
	}
}
