package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
