package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}
type Articles []article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title 1", Description: "Description test 1"},
		Article{Title: "Test Title 2", Description: "Description test 2"},
		Article{Title: "Test Title 3", Description: "Description test 3"},
		Article{Title: "Test Title 4", Description: "Description test 4"},
	}
	fmt.Println("End Points hit all the articles")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homePage endpoint hits")
}
func requestHandler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
}
func main() {
	requestHandler()
}