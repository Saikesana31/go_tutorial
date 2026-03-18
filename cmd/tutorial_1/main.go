package main // This is the main package

import "fmt"

// This is the main function
func main() {
	var Num int = 13245
	fmt.Println(Num)

	var floatNum float64 = 13245.1234567890
	fmt.Println(floatNum)

	var floatNum2 float32 = 13245.1234567890
	var intNum int = 13245
	var resutl float32 = floatNum2 + float32(intNum)
	fmt.Println(resutl)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var mystring string = "Hello, World!"
	fmt.Println(mystring)

	var myboolean bool = false
	fmt.Println(myboolean)

	// short declaration to create a variable
	myvar := "text"
	fmt.Println(myvar)

	// Constants[ once created, can't be changed ]
	const myconst string = "Can't be changed once created"
	fmt.Println(myconst)
}
