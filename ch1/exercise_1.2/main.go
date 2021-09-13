package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Printf("index: %d, value: %v\n", idx, arg)
	}
}
