package main

import "fmt"

/* Problem 1 */
func main() {
	var sum int
	upperBound := 1000
	for i := 0; i < upperBound; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

/*
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
 * The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */
