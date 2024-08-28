package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shinez1997/http/handlers"
)

func main() {
	l := log.New(os.Stdout, "preoduct-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
}
