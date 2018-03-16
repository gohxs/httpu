package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		extra := httpu.Param(r)

		log.Println("Extra:", extra)
	})

	log.Println("Listening at :8081")
	http.ListenAndServe(":8081", mux)
}
