package di

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

func Greet2(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
func Greet3(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	Greet3(w, "world")
}

func main() {
	// Greet3(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(Handler)))
}
