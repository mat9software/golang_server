package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// --------------------------------------------------
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("DEBUG:default handle")
		fmt.Fprintf(w, "Hi /")
	})

// ----------
 SessionHandlersInit()
 UserHandlersInit()

// ----------
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Print("Listening on port " + httpPort)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
