package main

import (
	"log"
	"net/http"
	"os"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "preoduct-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	http.ListenAndServe("localhost:8080", sm)
}
