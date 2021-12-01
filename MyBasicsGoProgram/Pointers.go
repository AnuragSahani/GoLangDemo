package main

import (
	"fmt"
)

func main() {

	num := 4
	//var n *int = &num
	sum(&num)
	fmt.Println(num)
	//fmt.Println(n)
}

func sum(num *int) {
	*num = 40
	fmt.Println(*num)
}
