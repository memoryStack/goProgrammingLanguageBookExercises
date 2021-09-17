package main

import "fmt"

func removeDuplicate(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

func main() {
	s := []string{"a", "b", "b", "b", "c"}

	s = removeDuplicate(s)
	fmt.Println(s)

}
