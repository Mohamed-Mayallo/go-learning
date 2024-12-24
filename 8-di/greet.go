package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, message string) {
	fmt.Fprint(w, message)
}

func main() {
	// in the application, use the Stdout, and in test use Buffer to be able to test
	// Greet(os.Stdout, "Hi")

	// another option in your application, you can use the http.ResponseWriter as a writer instead of the os.Stdout or Buffer
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "hello internet")
	})

	log.Fatal(http.ListenAndServe(":4040", nil))
}
