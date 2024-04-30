package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title string
	Director string
}

func main() {
	// Define the port
	PORT := "3000"
	fmt.Println("Server started on http://localhost:" + PORT)

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define a route handler for GET home
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// Define the data
		films := []Film {
			{Title: "Jurassic Park", Director: "Steven Spielberg"},
			{Title: "Matrix", Director: "Wachowski Brothers"},
			{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
		}

		// Define the template and return it along with the data
		response := template.Must(template.ParseFiles("index.html"))
		response.Execute(w, films)
	})

	// Listen on port 8000
	log.Fatal(http.ListenAndServe(":3000", nil))

	// TIP: log.Fatal is a wrapper around log.Print() followed by a call to os.Exit(1), it logs the error and exits the program if an error occurs
}

// Run the program with: `go run main.go`
// Or, install gomon (nodemon for go): `go install github.com/julesguesnon/gomon@latest`
// Add PATH to your .bashrc or .zshrc: `export PATH=$PATH:$HOME/go/bin`
// Run the program with: `gomon main.go`
