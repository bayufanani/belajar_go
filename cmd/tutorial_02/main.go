package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var myInt int = 32767
	myInt = myInt + 1
	fmt.Println(myInt)

	var myFloat float32 = 12345678.9
	fmt.Println(myFloat)

	var floatNum32 float32 = 10.2
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)

	var myString string = "hello world"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("γ"))

	var a rune = 'a'
	fmt.Println(a)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	var myString1 = "asd"
	fmt.Println(myString1)

	myString2 := "this is a string"
	fmt.Println(myString2)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// give error
	const myConst int = 3
	fmt.Println(myConst)
}
