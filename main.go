package main

import (
	"log"
	"net/http"
)

func main() {
	//creates fileserver
	fileserver := http.FileServer(http.Dir("static/"))

	//pass request to server
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fileserver.ServeHTTP(w, r)
		},
	)
	// start http server on port 8080
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
