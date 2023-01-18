// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Go on GKE!"))
    })
	log.Print("Running main func")
    http.ListenAndServe(":8080", nil)
}
