package main

import (
	"os"
	simpleserver "simpleserver/internal"
)

func main() {
	port := os.Args[1]
	simpleserver.Start(port)
}
