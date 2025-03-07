package main

import "fmt"

func main() {
	var intNum int16 = 4568
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 121213131.31
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 311
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 1
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello \nworld"
	fmt.Println(myString)

	var myString1 string = "Hello" + " " + "World"
	fmt.Println((myString1))

	fmt.Println(len(myString1))

	var myRune rune = 'a'
	fmt.Println(myRune) // ASCII value

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var myVar = "text"
	fmt.Println(myVar)

	const myConst string = "const value"
	fmt.Println(myConst)

}
