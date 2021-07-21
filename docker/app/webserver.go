// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// var (
// 	Client *sql.DB

// 	username = os.Getenv("MYSQL_USER")
// 	password = os.Getenv("MYSQL_PASSWORD")
// 	host     = os.Getenv("MYSQL_HOST")
// 	db       = os.Getenv("MYSQL_DB")
// )

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := "bork"
	// f, _ := os.Open("webserver.go")
	// defer f.Close()
	// io.Copy(w, f)
	fmt.Fprintf(w, "hello %s!\n", name)
}
