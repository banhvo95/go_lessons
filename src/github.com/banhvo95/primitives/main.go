package main

import (
	"fmt"
)

func main() {
	//Boolean type
	// 0 value for bool is false ( in case where you dont declare)
	// ------------------------------------------
	// fmt.Println("HELLO WORLD")

	// var n bool
	// fmt.Printf("%v, %T\n", n, n)

	// x := 1 == 1
	// y := 1 == 2

	// fmt.Println(x, y)

	// ------------------------------------------
	// Numeric Types

	// int, int8, int16, int32, int64
	// unsigned ints - uint16.. etc
	n := 42
	fmt.Printf("%v, %T\n", n, n)
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // integer division
	fmt.Println(a % b) // remainder

	//make sure types match int8/int32 == bad
	// bit operators &, |, ^, &^ and, or, exclusive or, and not

	c := 8 // 1000
	//bit shifting
	fmt.Println(c << 3) //100000 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) //0001 2^3 / 2^3 = 2^0

	d := 3.14
	d = 13.7e72
	d = 2.1e14
	fmt.Printf("%v, %T\n", d, d)

	var ComplexNum complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", ComplexNum, ComplexNum)
	//Text Type

	// strings can be treated like arrays
	//strings are alias in byte
	var TestString string = "HELLO WORLD"
	fmt.Printf("%v, %T\n", TestString[3], TestString[3])         //returns 76, uint8
	fmt.Printf("%v, %T\n", string(TestString[3]), TestString[3]) //returns L, uint8

	// strings are immutable
	// TestString[3] = "b" will fail

	// can concat strings TestString1 + TestString2

	// can also slice bytes
	by := []byte(TestString)
	fmt.Printf("%v, %T\n", by, by) //[72 69 76 76 79 32 87 79 82 76 68], []uint8
	// Useful for taking advantage of this

	// RUNE REPRESENTS UTF-32 [can be 32 bits, but doesnt have to be] UTF-8 is utf-32
	// true type alias, same thing as a int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r) // returns 97, int32
}
