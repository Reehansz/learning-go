package main

import "fmt"

func DemoOperator() {
	// Arithmetic Operators
	a := 10
	b := 5
	fmt.Println("\nArithmetic Operators:")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Assignment Operators
	fmt.Println("\nAssignment Operators:")
	c := a
	fmt.Println("c =", c)
	c += b
	fmt.Println("c += b ->", c)
	c -= b
	fmt.Println("c -= b ->", c)
	c *= b
	fmt.Println("c *= b ->", c)
	c /= b
	fmt.Println("c /= b ->", c)
	c %= b
	fmt.Println("c %= b ->", c)

	// Comparison Operators
	fmt.Println("\nComparison Operators:")
	fmt.Println("a == b ->", a == b) // Equal to
	fmt.Println("a != b ->", a != b) // Not equal to
	fmt.Println("a > b ->", a > b)   // Greater than
	fmt.Println("a < b ->", a < b)   // Less than
	fmt.Println("a >= b ->", a >= b) // Greater than or equal to
	fmt.Println("a <= b ->", a <= b) // Less than or equal to

	// Logical Operators
	fmt.Println("\nLogical Operators:")
	x := true
	y := false
	fmt.Println("x && y ->", x && y) // Logical AND
	fmt.Println("x || y ->", x || y) // Logical OR
	fmt.Println("!x ->", !x)         // Logical NOT

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("a & b =", a&b)   // Bitwise AND
	fmt.Println("a | b =", a|b)   // Bitwise OR
	fmt.Println("a ^ b =", a^b)   // Bitwise XOR
	fmt.Println("a &^ b =", a&^b) // Bit clear (AND NOT)
	fmt.Println("a << 1 =", a<<1) // Left shift
	fmt.Println("a >> 1 =", a>>1) // Right shift

	str1 := "Reehan"
	str2 := "Reehan"
	fmt.Println(str1 == str2)
}
