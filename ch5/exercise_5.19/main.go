package main

import "fmt"

func passByPanic(panicRecover func()) {
	defer panicRecover()
	a := 1
	panic(a)

}

func main() {
	var valueToBeReturned int
	recoverFunc := func() {
		valueToBeReturned = recover().(int)
	}
	passByPanic(recoverFunc)
	fmt.Println(valueToBeReturned)
}
