package main

import (
	"fmt"
	"io"
	"os"
)

// Greet accpets a person's name then prints a greeting
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
