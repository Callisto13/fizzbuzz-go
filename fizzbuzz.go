package fizzbuzz

func IsDivisibleByThree(num int) bool {
	return num%3 == 0
}

func IsDivisibleByFive(num int) bool {
	return num%5 == 0
}

func IsDivisibleByThreeAndFive(num int) bool {
	return num%15 == 0
}
