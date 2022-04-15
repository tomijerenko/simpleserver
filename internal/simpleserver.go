package simpleserver

import (
	"fmt"
	"log"
	"net/http"
)

type endpoint struct {
	hostname string
	port     string
}

func Start(port string) {
	e := &endpoint{
		hostname: "localhost",
		port:     port,
	}

	http.HandleFunc("/", e.base)
	http.HandleFunc("/left", e.left)
	http.HandleFunc("/right", e.right)

	fmt.Println("Server listening on " + e.hostname + ":" + e.port)
	log.Fatal(http.ListenAndServe(":"+e.port, nil))
}

func (e *endpoint) base(w http.ResponseWriter, r *http.Request) {
	out := e.hostname + ":" + e.port + "/"
	fmt.Println(out)
	fmt.Fprintf(w, out)
}

func (e *endpoint) left(w http.ResponseWriter, r *http.Request) {
	out := e.hostname + ":" + e.port + "/left"
	fmt.Println(out)
	fmt.Fprintf(w, out)
}

func (e *endpoint) right(w http.ResponseWriter, r *http.Request) {
	out := e.hostname + ":" + e.port + "/right"
	fmt.Println(out)
	fmt.Fprintf(w, out)
}
