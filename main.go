package main

import (
	"fmt"
	"net/http"
)

func main() {
	servMux := http.NewServeMux()
	servMux.Handle("/", http.FileServer(http.Dir(".")))
	server := http.Server{
		Addr:    ":8080",
		Handler: servMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
