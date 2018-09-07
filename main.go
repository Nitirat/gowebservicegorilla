package main

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/greets/{name}", greeting)
	fmt.Println("start server")
	http.ListenAndServe(":8080", mux)
}

func greeting(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the home page! %s", name)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the home page!")
}
