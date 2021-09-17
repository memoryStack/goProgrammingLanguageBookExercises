package main

import "fmt"

// TODO: how to achieve this for generic length arrays ??
func reverse(ptr *[5]int) {

	for i := 0; i < 5/2; i++ {
		ptr[i], ptr[5-i-1] = ptr[5-i-1], ptr[i]
	}

}

func main() {

	array := [5]int{0, 1, 2, 3, 4}

	reverse(&array)

	fmt.Println(array)

}
