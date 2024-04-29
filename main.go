package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Define the port
	PORT := "3000"
	fmt.Println("Server started on http://localhost:" + PORT)

	// Define a route handler
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World\n")
		io.WriteString(w, r.Method)
		// OUTPUT: Hello World GET
	})

	// Listen on port 8000
	log.Fatal(http.ListenAndServe(":3000", nil))
	

	// TIP: log.Fatal is a wrapper around log.Print() followed by a call to os.Exit(1), it logs the error and exits the program if an error occurs
}

// Run the program with: `go run main.go`
// Or, install gomon (nodemon for go): `go install github.com/julesguesnon/gomon@latest`
// Add PATH to your .bashrc or .zshrc: `export PATH=$PATH:$HOME/go/bin`
// Run the program with: `gomon main.go`
