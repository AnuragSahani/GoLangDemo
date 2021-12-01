//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//)
//
//type Article struct {
//	Title   string `json:"title,omitempty"`
//	Desc    string `json:"desc,omitempty"`
//	Content string `json:"content,omitempty"`
//}
//
//type Articles []Article
//
////@get when we type /articles
//
//func allArticales(w http.ResponseWriter, r *http.Request) {
//	articles := Articles{
//		Article{Title: "Test Title", Desc: "Test Description", Content: "I am content"},
//	}
//
//	fmt.Println("EndPoint Hit: All artical end point")
//	json.NewEncoder(w).Encode(articles)
//}
//
////@get when we type /
//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "This is Our Homepage")
//}
//
////@Entry Point
//func handleRequest() {
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/articles", allArticales)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}
//
//func main() {
//	handleRequest()
//	// handleRequest is a entry Point
//	// homepage is post method
//	// allArticles is get method that send some json response
//
//	// How we check their responses through POSTMAN
//}
