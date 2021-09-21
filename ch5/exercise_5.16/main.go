package main

import (
	"fmt"
	"strings"
)

func variadicStrJoin(sep string, vals ...string) string {

	return strings.Join(vals, sep)

}

func main() {
	// strings.Join(a, sep)
	fmt.Println(variadicStrJoin(",", "anuj", "rao"))
}
