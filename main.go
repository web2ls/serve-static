package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	const PORT = ":8080"
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Println("Server has been started on port " + PORT)
	http.ListenAndServe(PORT, nil)
}
