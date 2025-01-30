package main

import (
	"fmt"
)

func main() {
	DemoBasic()
	DemoControlStructure()
	DemoFunctions()
	DemoOperator()

	array := InitArray()
	fmt.Println(array)

	string1 := InitString()
	fmt.Println(string1)

	numbers := InitArraySlice()

	fmt.Printf("The type of the object is %T \n", numbers)
}
