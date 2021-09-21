package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$foo", f("foo"))
}

func main() {

	// TODO: do it with all sorts of io/out options to practice the io/out libraries in GO
	bytes, _ := ioutil.ReadAll(os.Stdin) // ignoring the error

	inputStr := string(bytes)

	f := func(s string) string {
		return s + "_" + "coverted"
	}
	fmt.Println(expand(inputStr, f))
}
