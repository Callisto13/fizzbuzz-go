package fizzbuzz

import "fmt"

func IsDivisibleBy(number, divisor int) bool {
	return number%divisor == 0
}

func Says(number int) string {
	switch {
	case IsDivisibleBy(number, 15):
		return "fizzbuzz"
	case IsDivisibleBy(number, 3):
		return "fizz"
	case IsDivisibleBy(number, 5):
		return "buzz"
	default:
		return fmt.Sprintf("%d", number)
	}
}
