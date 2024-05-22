package main

import (
	"fmt"
	"log"
	"net/http"
)

// A func to Handle /hello route -> A simple go function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Printing from Hello Handler")
}

// A simple handler to handle input data through form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)

}
func main() {
	// A simple route to Serve a file to a web server
	fileHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandle)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
