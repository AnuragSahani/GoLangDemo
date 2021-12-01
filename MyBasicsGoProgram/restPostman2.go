package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Content string `json:"content,omitempty"`
}

type Articles []Article

func allArticalesPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Our allArticalPost method, I don't have any data")
}

//@get when we type /articles

func allArticales(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "I am content"},
		Article{Title: "Sapna", Desc: "Test for sapna Description", Content: "I am sapna content"},
	}

	fmt.Println("bhai user ne allArtical method hit kiya hai")
	json.NewEncoder(w).Encode(articles)
}

//@get when we type /
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Our Homepage")
}

//@Entry Point
func handleRequest() {

	// For Mux Usage

	myRouter := mux.NewRouter().StrictSlash(true)

	// Change http through myRouter jab / hota hai to root

	myRouter.HandleFunc("/", homePage)

	//Mux help to identifies which one is GET method and POST METHOD
	myRouter.HandleFunc("/articles", allArticales).Methods("GET")
	myRouter.HandleFunc("/articles", allArticalesPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
	// handleRequest is a entry Point
	// homepage is post method
	// allArticles is get method that send some json response

	// How we check their responses through POSTMAN
}
