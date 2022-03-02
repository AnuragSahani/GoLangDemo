package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Students struct {
	id        string `validate:"omitempty,uuid"`
	firstName string `validate:"required"`
	email     string `validate:"required,email"`
	userName  string `validate:"required,email"`
	password  string `validate:"required,gte=8"`
	college   string `validate:"uppercase"`
}

func main() {

	res := Students{
		id:        "1",
		firstName: "Anurag",
		email:     "anurag0123",
		userName:  "anuragsahani0123@gmail.com",
		password:  "1234567",
		college:   "BIT",
	}
	validate := validator.New()
	err := validate.Struct(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Printf("%v", res)

}
