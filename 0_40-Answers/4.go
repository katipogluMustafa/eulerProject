package main

import (
	"fmt"
	"math"
)

/* Problem 4 */
/*
 * Answer is almost true
 * Debug the case when you shrink the number if there is a number zero at the beginning of the number
 *
 */
func main() {
	fmt.Println(findTheAnswer4())
}
func findTheAnswer4() int {
	var max int
	for i := 999; i >= 800; i-- {
		for j := 999; j >= 800; j-- {
			product := i * j
			if isPalindrome(product) {
				if product > max {
					max = product
				}
			}
		}
	}
	return max
}
func isPalindrome(n int) bool {
	boolly := true
	digit := numOfDigits(n)
	for boolly == true && digit >= 1 {
		if !isFirstAndLastEqual(n, digit) {
			boolly = false
		}
		n = (n % int(math.Pow10(digit-1))) / 10
		digit -= 2

	}
	return boolly
}
func isFirstAndLastEqual(n int, digit int) bool {
	if n%10 == n/int(math.Pow10(digit-1)) {
		return true
	}
	return false
}
func numOfDigits(n int) int {
	var digit int

	for n > 0 {
		digit++
		n /= 10
	}

	return digit
}

/*
 * A palindromic number reads the same both ways.
 * The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */
