package main

import (
	"fmt"
)

/*
 *
 * WRONG ANSWER, Look This One Later Again
 *
 */
/* Problem 3 */
func main() {
	fmt.Println(largestPrime(497))
}
func largestPrime(n int) int {
	max := -1
	var i int
	half := n / 2
	if half%2 == 1 {
		i = half
	} else {
		i = half + 1
	}
	for i > 1 {
		if isPrime(i) {
			max = i
			return max
		}
		i -= 2
	}
	return max
}
func isPrime(p int) bool {
	if p <= 1 {
		return false
	}
	if largestPrime(p) == -1 {
		return true
	}
	return false
}

/*
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 */
