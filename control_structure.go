package main

import "fmt"

func DemoControlStructure() {
	//if - else
	age := 18
	if age >= 18 {
		fmt.Println("\nYou are an adult")
	} else {
		fmt.Println("\nYou are not an adult")
	}

	//loops (there is no while loop in go, but we can use for as a while loop)
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//for as a while loop
	cnt := 5
	for cnt <= 10 {
		fmt.Println(cnt)
		cnt++
	}

	// Iterating over a slice using range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Its a normal day")
	}
}
