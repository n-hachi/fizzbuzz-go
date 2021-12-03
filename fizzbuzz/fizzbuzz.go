package fizzbuzz

import "fmt"

// Fizzbuzz takes one postive value and return string "Fizz", "Buzz", "FizzBuzz" or number.
func Fizzbuzz(n uint64) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
