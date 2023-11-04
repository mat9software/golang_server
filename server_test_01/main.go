package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	log.Print("Listening on port " + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))

}
