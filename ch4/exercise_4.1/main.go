// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"crypto/sha512"
	"flag"
	"fmt"
)

func decode256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func decode384() {
	c1 := sha512.Sum384([]byte("x"))
	c2 := sha512.Sum384([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func decode512() {
	c1 := sha512.Sum512([]byte("x"))
	c2 := sha512.Sum512([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func main() {

	shaFlag := flag.Int("sha", 256, "use specific sha func, default -> sha256, 1 -> sha384, 2 -> 512")
	flag.Parse()

	switch *shaFlag {
	case 256:
		decode256()
		break
	case 384:
		decode384()
		break
	case 512:
		decode512()
	default:
		fmt.Println("invalid sha variant")
	}
}

//!-
