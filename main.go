package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login attempt")
	w.Write([]byte("Login OK"))
}

func main() {
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
