// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func max(a int, vals ...int) int {
	maxVal := a
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func min(a int, vals ...int) int {
	minVal := a
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

//!-

func main() {

	fmt.Println(max(1, 2, 3, 4, -1))
	fmt.Println(min(1, 2, -3, 4, -1))

	// TODO: how to handle these kind of cases ??
	// if i want to pass an array then i have to pass it like below which
	// is not easy at all
	values := []int{1, 2, 3}
	fmt.Println(min(values[0], values[1:]...))
}
