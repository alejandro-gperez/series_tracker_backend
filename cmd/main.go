package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("polo"))
	})

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", nil)
}
