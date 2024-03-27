package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		_, err := fmt.Fprintf(w, "ParseForm() err: %v", err)
		if err != nil {
			return
		}
		return
	}
	_, err := fmt.Fprintf(w, "POST request successful")
	if err != nil {
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	_, err = fmt.Fprintf(w, "Name = %s\n", name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "Address = %s\n", address)
	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(w, "hello!")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
