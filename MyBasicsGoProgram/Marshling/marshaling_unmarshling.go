package main

import (
	"encoding/json"
	"fmt"
)

type securaa struct {
	ID          int
	Name        string
	Departments []string
}

type securaa_advance struct {
	ID          int
	Name        string
	Departments string
}

func main() {

	nasccom := securaa{
		ID:          1,
		Name:        "anurag",
		Departments: []string{"testing", "BackEnd", "UI"},
	}

	temp, err := json.Marshal(nasccom)
	if err != nil {
		fmt.Println("error:", err)
	}
	pp := string(temp)
	//fmt.Println(nasccom)
	fmt.Println(pp)
	//os.Stdout.Write(temp)

	//err = json.Unmarshal(temp, nasccom)
	//if err == nil {
	//	fmt.Println("error in unMarshaling ", err)
	//}
	//os.Stdout.Write(temp)

	//............UnMarshaling..................

	var jsonByteDataForFilling = []byte(`[
{"Id": 33, "Name":"Neeraj Tiwari","Departments" :"UI"}
]`)

	var sec []securaa_advance
	err = json.Unmarshal(jsonByteDataForFilling, &sec)
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Printf("%+v", sec)

}
