package main

import "fmt"

func main() {
	var myString = "résumé"
	var indexedString = myString[0]
	fmt.Printf("%v, %T\n", indexedString, indexedString)
}
