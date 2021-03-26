package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johnmwood/bgtopfive/api/youtube"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Board Game Top Five")
	fmt.Println("Hit endpoint: homepage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	youtube.SetYoutubeConfig()
	handleRequests()
}
