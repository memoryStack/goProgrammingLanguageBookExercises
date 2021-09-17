package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	solution inspiration
		https://www.geeksforgeeks.org/array-rotation/
	TODO: i would like to find out some day why this gcd thing works. and how can i think like this.
					maybe i have to learn number theory better
*/

func getNextIdx(idx, pos, len int) int {
	if idx <= pos {
		return len - 1 - (pos - idx)
	}
	return idx - pos - 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// put [0...pos] elements in the end of the array
func rotate(s []int, pos int) {

	if len(s) < 2 || (pos < 0 || pos > len(s)) {
		return
	}

	actualMoves := gcd(pos+1, len(s))

	for i := 0; i < actualMoves; i++ {
		// value to be placed

		currentIndx := i
		value := s[currentIndx]
		for {
			nextIndex := getNextIdx(currentIndx, pos, len(s))
			value, s[nextIndex] = s[nextIndex], value
			currentIndx = nextIndex
			if currentIndx == i {
				break
			}
		}

	}

}

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6}

	pos := os.Args[1]

	posInt, _ := strconv.ParseInt(pos, 10, 32)

	rotate(slice, int(posInt))
	fmt.Println(slice)
}
