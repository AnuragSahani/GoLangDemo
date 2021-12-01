// fun show()
// go show()
//

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Go Routine karne waale hai")

	go routine("no blocking 1")
	routine("sequential")
	go routine("no blocking 2")

	fmt.Scanln()
}

func sum(num *int) {
	*num = 40
	fmt.Println(*num)
}

func routine(s string) {
	for i := 0; i < 50; i++ {
		fmt.Println(i, s)
	}
}
