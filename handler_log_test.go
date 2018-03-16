package httpu_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gohxs/httpu"
)

func ExampleLogHandler() {
	llog := log.New(os.Stdout, "[main] ", 0)

	mux := http.NewServeMux()
	mux.HandleFunc("/", httpu.LogHandler(llog, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Will log")
	}))
	s := httptest.NewServer(mux)
	http.Get(s.URL)
	// Output:
	// Will log
	// [main] (127.0.0.1) GET / - [200 OK]
}

func ExampleLogHandlerNotFound() {
	llog := log.New(os.Stdout, "[main] ", 0)

	mux := http.NewServeMux()
	mux.HandleFunc("/test/", httpu.LogHandler(llog, func(w http.ResponseWriter, r *http.Request) {
		if len(httpu.Params(r)) > 1 {
			httpu.WriteStatus(w, http.StatusExpectationFailed)
			return
		}
		httpu.WriteStatus(w, http.StatusNotFound)
	}))
	s := httptest.NewServer(mux)
	http.Get(s.URL + "/test")
	http.Get(s.URL + "/test/100/12")
	http.Get(s.URL + "/tes") // will not output
	// Output:
	// [main] (127.0.0.1) GET /test/ - [404 Not Found]
	// [main] (127.0.0.1) GET /test/100/12 - [417 Expectation Failed]
}
