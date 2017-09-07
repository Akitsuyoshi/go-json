package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is a struct for all articles
type Article struct {
	ID      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllAritcles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{ID: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("endpoint hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintf(w, "Key: "+key)
}

//homepage handle all requests to our root url
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welocome to the homepage")
	fmt.Println("endpoint hit: homepage")
}

//handleRequests match the url path with a defined function
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article/{key}", returnSingleArticle)
	myRouter.HandleFunc("/all", returnAllAritcles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
