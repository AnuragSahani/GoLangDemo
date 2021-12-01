package main

import (
	"fmt"
)

const (
	z = iota
	y = iota
	x
	w = iota
)

func main() {

	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(w)

	//Map <String,int> emp = new HashMap<>();

	emp := make(map[string]int) //declaration

	// Intialization
	emp := map[string]int{
		"Sapna Singh": 25000,
		"Kajal":       30000,
		"Sappu":       0,
	}

	fmt.Println(emp)

	find, ok := emp["Kajal"]
	fmt.Println(find, ok)

	a, ok := emp["Sappu"]

	if ok == true {
		fmt.Println(a, " Value is avilabe")
	} else {
		fmt.Println(a, " Value is Not avilabe")
	}

	student := map[int]string{
		40:  "Sapna",
		903: "anurag",
	}
	fmt.Println(student)

	fmt.Println(len(student))

	fmt.Println(student[903])

	student[40] = "Sapna Singh"

	student[901] = "AKM"
	//fmt.Println(student)

	delete(student, 901)
	fmt.Println(student)

	bitStudent := student
	fmt.Println(bitStudent)

	bitStudent[900] = "Kajal Singh"

	fmt.Println(student)
	fmt.Println(bitStudent)

	for i := 1; i <= 10; i++ {
		fmt.Println(i)

	}

	var i interface{} = 5.3
	cc := 5 + 3i
	fmt.Printf("%v %T \n", cc, cc)

	var a int = 5
	fmt.Println(a)
	a = 8
	fmt.Println(a)
	const b int = 8
	fmt.Println(b)

	i := 9
	switch i {
	case 1:
		fmt.Println("I am no. 1")
	case 2:
		fmt.Println("I am no. 2")
	case 3, 9:
		fmt.Println("I am no. 3")
	case 4:
		fmt.Println("I am no. 4")
	case 5:
		fmt.Println("I am no. 5")
		fallthrough
	case 6:
		fmt.Println("I am no. 6")
	default:
		fmt.Println("I am no. defaulter")

	case int:
		fmt.Println("It's a Int")
	case string:
		fmt.Println("It's a string")
	case float64:
		fmt.Println("It's a float64")
	default:
		fmt.Println("Anurag")
}

func Show(a, b int) int {
	return a + b
}

func sum(a ...int) {

	res := 0
	for _, n := range a {
		res += n
	}
	fmt.Println(res)
}
