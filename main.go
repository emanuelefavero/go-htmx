package main

import "fmt"

func main() {
	hello := "Hello" // := infers the type of the variable
	fmt.Println(hello)
}

// Run the program with: `go run main.go`
// Or, install gomon (nodemon for go): `go install github.com/julesguesnon/gomon@latest`
// Add PATH to your .bashrc or .zshrc: `export PATH=$PATH:$HOME/go/bin`
// Run the program with: `gomon main.go`
