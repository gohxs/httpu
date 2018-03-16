package main

import (
	"log"
	"net/http"

	"github.com/gohxs/httpu/chain"
)

func main() {
	mux := http.NewServeMux()

	c := chain.New(httpu.ChainLogger("main"))

	mux.HandleFunc("/", c.Build(httpu.CatchAllHandler(func(w http.ResponseWriter, r *http.Request) {
		param := httpu.Param(r)
		log.Println("Param is:", param)
		if param[0] == "hello" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Catching all")
	})))
	log.Println("Listening at :8080")
	http.ListenAndServe(":8080", mux)
}
