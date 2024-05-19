package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//  mod commands is expensive!!

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>hi</h2>"))
	// fmt.Println(r) // request
	// fmt.Println(w)
}

// go mod build
// go mod vendor
// go mod tidy
