package main

import (
	"fmt"
	"strconv"
)

// Global variable (Pascal Case)
var An int = 12

// Package level variable (Camale case)
var stringCase string = "String Camale Case Package level"

// Shadowing variable
var shad string = " I am Local Shadow variable"

func main() {

	// fmt.Println("Hello world")

	// // 1st method
	// var a int // declaration
	// a = 55    // Initalization
	// fmt.Println(a)

	// // 2nd method declare in one line
	// var b int = 44
	// fmt.Println(b)

	// // 3rd Declaration without giving DataType

	// var str = "Anurag"
	// fmt.Println(str)

	// //4th method ios very useful

	// d := "AnuragSahani"
	// e := 100
	// f := 14.22
	// g := 'a'

	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)
	// fmt.Println(g)

	// fmt.Println(An)
	// fmt.Println(stringCase)

	// // Shadowing variable
	// shad := " I am Local Shadow variable"
	// fmt.Println(shad)

	// // If we want to know the data type of varible
	// dataT := 100
	// fmt.Printf("%d, %T", dataT, dataT)
	// fmt.Println()

	// Type Conversion
	var z int32 = 100
	var y int64 = int64(z)
	fmt.Printf("%d , %T", y, y)
	fmt.Println()

	x := 15.55514
	var st int = int(x)
	fmt.Printf("%v , %T", st, st)
	fmt.Println()

	r := 65
	var strI string = strconv.Itoa(r)
	fmt.Printf("%v %T", strI, strI)
	fmt.Println()

	// If you want to print the ASCII value of any int value
	var d string = string(r)
	fmt.Printf("ASCII valur of variable r 65 = %v, %T", d, d)
	fmt.Println()

	// If you want to convert string to int
	stI1 := "2"
	a, err := strconv.Atoi(stI1)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v %T", a, a)
	fmt.Println()

	// Complax data type representation like var com complex128 = 5 +3i
	com := 5i
	fmt.Printf("%v %T", com, com)
	fmt.Println()

	//Array

	var arr [5]int = [5]int{1, 2, 1, 5, 4}
	fmt.Println(arr)

	var arr1 = [...]int{1, 2, 3, 4, 6, 5, 5, 6, 6, 6, 6, 6, 0}
	fmt.Println(arr1)

}
