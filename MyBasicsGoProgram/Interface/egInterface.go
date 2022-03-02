package main

import "fmt"

const (
	a   = 1
	b   = 4
	str = "aaaaaa"
)

func main() {

	var obj Vehical = Bike{
		name:  "Honda",
		color: "Red",
		price: 8800.2,
	}
	obj.ShowDetails()
	obj.Showname()
	fmt.Println(a)

}

type Vehical interface {
	ShowDetails()
	Showname()
}

type Bike struct {
	name, color string
	price       float64
}

func (bike Bike) Showname() {
	fmt.Println("Bike name ::==", bike.name)
}

func (bike Bike) ShowDetails() {
	fmt.Println("Bike Name : =", bike.name)
	fmt.Println("Bike color : =", bike.color)
	fmt.Println("Bike price : =", bike.price)
}

//TODO: remove Vehical interface and see
