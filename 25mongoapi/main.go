package main

import (
	"fmt"
	"log"
	"net/http"

	// Corrected import statement
	"github.com/hiteshchoudhary/mongoapi/router"
)

// main is the entry point for the MongoDB API server. It initializes the router
// and starts the HTTP server on port 4000, logging any fatal errors encountered
// during execution.

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
