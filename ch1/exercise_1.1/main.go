package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("executing command is: \n %v\n", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
