package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func DemoFunctions() {
	var a int = 5
	var b int = 6
	fmt.Print("\nEnter the value of a : \t")
	//fmt.Scan(&a)
	fmt.Print("Enter the value of b : \t")
	//fmt.Scan(&b)
	result := add(a, b)
	fmt.Println("the output is ", result)
	res_div, err := divide(float32(a), float32(b))
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Output is : ", res_div)
	}
}
