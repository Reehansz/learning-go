package main

import "fmt"

// Slices are dynamic and flexible views into arrays. Unlike arrays, slices can be resized.

func InitArraySlice() []int {
	//initialize the slice with initial values
	slice := []int{10, 20, 30}
	fmt.Println("Slice : ", slice)

	fmt.Println("Appending slice now")
	slice = append(slice, 40)
	slice = append(slice, 50, 60) // multiple values
	fmt.Println("This is the appended slice : ", slice)
	//appending using another slice
	fmt.Println("This is the appended slice with another slice : ")
	slice2 := []int{70, 80, 90}
	slice = append(slice, slice2...)
	fmt.Println("This is the appended slice : ", slice)

	//creating a sub slice
	subslice := slice[1:4]
	fmt.Println("The subslice is : ", subslice)
	return slice
}
