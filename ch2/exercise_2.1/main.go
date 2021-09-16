package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goProgrammingLanguageBook/ch2/exercise_2.1/tempconv"
)

func main() {

	tempInKelvin := os.Args[1]

	tempInKelvinFloat, err := strconv.ParseFloat(tempInKelvin, 64)
	if err != nil {
		return
	}

	fmt.Println("in celcius: ", tempconv.KToC(tempconv.Kelvin(tempInKelvinFloat)))
	fmt.Println("in farenheight: ", tempconv.KToF(tempconv.Kelvin(tempInKelvinFloat)))
}
