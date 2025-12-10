package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)                // Register "/" route
	http.HandleFunc("/about", aboutHandler)          // Register "/about" route
	http.HandleFunc("/api/user", userHandler)        // Register "/api/user route"
	http.HandleFunc("/greet", greetHandler)          // Register "/greet route"
	http.HandleFunc("/login-page", loginPageHandler) // Shows the form
	http.HandleFunc("/login", loginHandler)          // Processes the form

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Start server on port 8080

}

func handleName(w http.ResponseWriter, r *http.Request) {
	// w = where you write your response
	// r = incoming request data
	fmt.Fprintf(w, "Your response text here")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
	fmt.Println(r)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
	fmt.Println(r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := map[string]string{
		"name": "testName",
		"age":  "testAge",
		"city": "testCity",
	}

	json.NewEncoder(w).Encode(user)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Hello, %s!", name)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<html>
          <body>
              <h1>Login Page</h1>
              <form action="/login" method="POST">
                  <input type="text" name="username" placeholder="Username"><br><br>
                  <input type="password" name="password" placeholder="Password"><br><br>
                  <button type="submit">Login</button>
              </form>
          </body>
          </html>
	`
	fmt.Fprintf(w, html)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Fprintf(w, "Recieved - Username: %s, Password: %s", username, password)
	} else {
		fmt.Fprintf(w, "Please use POST method")
	}
}
