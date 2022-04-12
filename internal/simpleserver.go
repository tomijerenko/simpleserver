package simpleserver

import (
	"fmt"
	"log"
	"net/http"
)

func Start(address string) {
	http.HandleFunc("/left", left)

	http.HandleFunc("/right", right)

	fmt.Println("Server listening on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func left(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LEFT was triggered")
	fmt.Fprintf(w, "LEFT")
}

func right(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RIGHT was triggered")
	fmt.Fprintf(w, "RIGHT")
}
