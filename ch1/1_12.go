package ch1

import (
	"log"
	"net/http"
	"strconv"
)

// Ex12 - Server1 is a minimal "echo" server.
func Ex12() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cyclesStrSlice := r.URL.Query()["cycles"]
		var cyclesInt float64
		cyclesInt = 5
		if cyclesStrSlice != nil {
			cyclesInt, _ = strconv.ParseFloat(cyclesStrSlice[0], 64)
		}
		Lissajous(w, cyclesInt)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
