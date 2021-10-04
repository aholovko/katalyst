package fizzbuzz

import "fmt"

func Convert(n int) string {
	var result string

	if n%3 == 0 {
		result = "Fizz"
	}

	if n%5 == 0 {
		result += "Buzz"
	}

	if result == "" {
		result = fmt.Sprintf("%d", n)
	}

	return result
}
