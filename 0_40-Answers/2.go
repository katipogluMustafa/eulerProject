package main

import (
	"fmt"
	"math"
)

/* Problem 2 */
func main() {
	var sum int
	var i int
	for fibo(i) < int(4*math.Pow(10, 6)) {
		temp := fibo(i)
		if temp%2 == 0 {
			sum += temp
		}
		i++
	}
	fmt.Println("Sum: ", sum)
}
func fibo(n int) int {
	i := [2][2]int{{1, 1}, {1, 0}}
	output := exponent(i, n)
	return output[0][1]
}
func multiply(first [2][2]int, second [2][2]int) [2][2]int {
	var output [2][2]int
	output[0][0] = first[0][0]*second[0][0] + first[0][1]*second[1][0]
	output[0][1] = first[0][0]*second[0][1] + first[0][1]*second[1][1]
	output[1][0] = first[1][0]*second[0][0] + first[1][1]*second[1][0]
	output[1][1] = first[1][0]*second[0][1] + first[1][1]*second[1][1]

	return output
}

func exponent(arr [2][2]int, n int) [2][2]int {
	if n%2 == 0 {
		return evenExponent(arr, n)
	} else {
		return oddExponent(arr, n)
	}

}

func oddExponent(first [2][2]int, n int) [2][2]int {
	return multiply(first, exponent(first, n-1))
}

func evenExponent(arr [2][2]int, n int) [2][2]int {
	if n == 0 {
		var out = [2][2]int{{1, 0}, {0, 1}}
		return out
	}
	out := exponent(arr, n/2)
	return multiply(out, out)
}

/*
 * Each new term in the Fibonacci sequence is generated by adding the previous two terms.
 * By starting with 1 and 2, the first 10 terms will be:
 *	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 * By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
 */
