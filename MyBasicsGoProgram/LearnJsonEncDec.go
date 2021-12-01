package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// variable declaration in struct
type Data struct {
	Username string
	Password string
	Email    string
}

// struct jo ke help kega container ke trah quki mutiple data store karaana tha that's why we using slicing
type Userdb struct {
	Users []Data
	Type  string
}

func main() {
	//createJsonFile()
	readJsonFile()
}

func createJsonFile() {
	users := []Data{
		{Username: "Anurag Sahani", Password: "change me", Email: "anuragsahani1@gmail.com"},
		{Username: "Avinash", Password: "must change me", Email: "avi1@gmail.com"},
		{Username: "Sapna", Password: "please change me", Email: "sappuss40@gmail.com"},
		{Username: "kajal", Password: "please don't change me", Email: "kajal40@gmail.com"},
	}

	db := Userdb{Users: users, Type: "friends"}

	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)

	f, err := os.Create("sappu_friends.db.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}

func readJsonFile() {
	db := Userdb{}

	f, err := os.Open("sappu_friends.db.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.Decode(&db)

	fmt.Println(db)

}
