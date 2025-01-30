package main

import "fmt"

func DemoBasic() {
	var a int = 5 // we can either use this type to declare variables or like the below shorthand :=
	d := "This is me Reehan"
	fmt.Println("\nHello world!", 23, "this is test")
	fmt.Println("Hey this is Reehan, How are you?")
	fmt.Println("Hello", a, "again")
	fmt.Println(d)

	// Basic Types: int, float64, string, bool
	// Composite Types: Arrays, Slices, Maps, Structs
	var important bool = true // important := true
	fmt.Println(important)

	//Using format specifiers to print
	a1 := 42
	b := 3.14159
	c := "Hello, Go!"
	d1 := true

	fmt.Printf("Integer: %d, Binary: %b, Hex: %x\n", a1, a1, a1)
	fmt.Printf("Float: %f, Scientific: %e\n", b, b)
	fmt.Printf("String: %s, Quoted String: %q\n", c, c)
	fmt.Printf("Boolean: %t\n", d1)
	fmt.Printf("Type of a: %T\n", a)

}
