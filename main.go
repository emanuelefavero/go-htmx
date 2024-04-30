package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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

	// Define the data
	films := []Film {
		{Title: "Jurassic Park", Director: "Steven Spielberg"},
		{Title: "Matrix", Director: "Wachowski Brothers"},
		{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
	}

	// * home route
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Define the template and return it along with the data
			response := template.Must(template.ParseFiles("index.html"))
			response.Execute(w, films)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // 405
		}
	})

	// * /add-film route
	http.HandleFunc("/add-film", func (w http.ResponseWriter, r *http.Request) {
		// Check if the request is an HTMX request (HX-Request header is set to true if it is)
		// log.Print(r.Header.Get("HX-Request"))

		switch r.Method {
		case "POST":
			// Simulate 0.5 second delay to show spinner
			time.Sleep(500 * time.Millisecond)

			// Parse the form data
			title := r.PostFormValue("title")
			director := r.PostFormValue("director")

			// Append the new film to the films slice
			films = append(films, Film{Title: title, Director: director})

			// Return the new film as a list item (before the user refreshes the home page, the new film will be added to the list without a page refresh, thanks to htmx and go templates)
			htmlString := fmt.Sprintf("<li class='list-group-item'>%s - Directed by %s</li>", title, director)
			response, _ := template.New("t").Parse(htmlString)
			response.Execute(w, nil)

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // 405
		}
	})

	// Listen on port 3000
	log.Fatal(http.ListenAndServe(":3000", nil))

	// TIP: log.Fatal is a wrapper around log.Print() followed by a call to os.Exit(1), it logs the error and exits the program if an error occurs
}

// Run the program with: `go run main.go`
// Or, install gomon (nodemon for go): `go install github.com/julesguesnon/gomon@latest`
// Add PATH to your .bashrc or .zshrc: `export PATH=$PATH:$HOME/go/bin`
// Run the program with: `gomon main.go`
