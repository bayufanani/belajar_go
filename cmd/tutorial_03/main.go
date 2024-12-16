package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "hello world"
	printMe(printValue)

	var nominator int = 3
	var denominator int = 0
	var result, remainder, err = intDivision(nominator, denominator)
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("result of int is %v", result)
	// } else {
	// 	fmt.Printf("Hasilnya adala: %v dengan sisa %v\n", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Printf("%s\n", err.Error())
	case remainder == 0:
		fmt.Printf("result of int is %v", result)
	default:
		fmt.Printf("Hasilnya adala: %v dengan sisa %v\n", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("pembagi tidak boleh 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
