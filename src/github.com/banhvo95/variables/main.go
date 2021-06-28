package main

import (
	"fmt"
	"strconv"
)

// cannot use := syntax globally
//"lower case" means it is scoped to the package it's written in
// "Upper case" letter means the compiler will expose it to the outside world
// "block scoped" if declared within {} it can only be accessed there

var global int = 16
var I int
var i int

func main() {
	fmt.Println("Hello Go")

	//All variables have to be used
	//if not it becomes a compile time error "x declared and not used"

	// How to declare different variables
	// Default value has a "0" value
	var i int
	fmt.Println(i)

	// Can do in line set
	var y int = 64
	fmt.Println(y)

	// Automatically assumes the type
	var x = 24
	fmt.Println(x)

	//Shorter hand way of doing it
	p := 13
	fmt.Println(p)

	fmt.Println(global)

	// Can wrap things in a "var" block
	var (
		name string = "John Doe"
		age  int    = 25
		pet  string = "Marlo"
	)
	fmt.Println(name, age, pet)

	// here global is shadowed in the inner scope
	var global int = 20
	fmt.Println(global)

	// NAMING CONVENTION
	// "lower case" means it is scoped to the package it's written in
	// "Upper case" letter means the compiler will expose it to the outside world
	// ---------------------------------------
	// variable names should be declared with lifespan of it in though
	// short name variables (for loops, etc)
	// long name variables if referenced more often and longer life span
	// for package level variables, you can make it longer
	// Acronyms - keep it as all uppercase for readability
	// var theURL string = "https...."
	// ---------------------------------------
	// Pascalcase or CamelCase, no underscore

	// ---------------------------------------
	// converting variables
	var i2 int = 42
	fmt.Printf("%v, %T\n", i2, i2)

	var j2 float32
	fmt.Printf("%v, %T\n", j2, j2)

	j2 = float32(i2)
	fmt.Printf("%v, %T\n", j2, j2)

	// Can cause loss of information (float -> int)
	// need to explicitly convert or itll cause an error
	//j2=i2 is bad

	//strings in go is an alias for string of bytes
	//string(32) will give 32 unicode

	//need strconv package for number to strings
	var j3 string
	j3 = strconv.Itoa(i2)
	fmt.Printf("%v, %T\n", j3, j3)

}
