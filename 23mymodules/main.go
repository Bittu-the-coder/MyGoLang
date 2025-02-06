package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in go lang")
	greater()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greater() {
	fmt.Println("Hey there mod in go lang")

}

func serverHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server Home")
	w.Write([]byte("<h1>Welcome to go lang series on youtube.</h1>"))

}
