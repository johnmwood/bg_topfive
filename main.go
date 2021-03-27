package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/johnmwood/bgtopfive/youtubeapi"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Board Game Top Five")
	fmt.Println("Hit endpoint: homepage")
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	youtubeapi.SetConfig()
	HandleRequests()
}
