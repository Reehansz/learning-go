package main

import (
	"fmt"

	"learning-packages/mathutil"
	"learning-packages/stringutil"
)

func main() {
	a := 5
	result := mathutil.Factorial(a)
	fmt.Printf("The factorial of %v is %v\n", a, result)

	var str string
	fmt.Println("Enter the string to reverse : ")
	fmt.Scan(&str)
	fmt.Println(stringutil.RevString(str))

	//errors in go
	fmt.Println("Enter the number to find the square root of :")
	var num float64
	fmt.Scan(&num)
	sqrt_num, err := mathutil.Sqrt(num)
	if err != nil {
		fmt.Println("There is an error in the number ", err)
	} else {
		fmt.Println("The square-root of the number is ", sqrt_num)
	}

}
