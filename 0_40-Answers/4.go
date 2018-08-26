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
	fmt.Println("6: ", isPalindrome(6))
	fmt.Println("36: ", isPalindrome(36))
	fmt.Println("33: ", isPalindrome(33))
	fmt.Println("363: ", isPalindrome(363))
	fmt.Println("364: ", isPalindrome(364))
	fmt.Println("3663: ", isPalindrome(3663))
	fmt.Println("3664: ", isPalindrome(3664))
	fmt.Println("36463: ", isPalindrome(36463))
	fmt.Println("36464: ", isPalindrome(36464))
	fmt.Println("364463: ", isPalindrome(364463))
	fmt.Println("364464: ", isPalindrome(364464))
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
