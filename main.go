package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/domainr/whois"
)

func main() {
	// fileServer for handling static directory
	fileServer := http.FileServer(http.Dir("static/"))

	// handling requests
	// Pass request to server
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		},
	)

	// /whois POST request
	// Client will send url-encoded request
	http.HandleFunc("/whois", func(w http.ResponseWriter, r *http.Request) {
		// Verify that request is POST request
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract data to perform function on
		data := r.PostFormValue("data")

		// whois query
		result, err := whoisQuery(data)

		// Return a encoded json
		if err != nil {
			jsonResponse(w, Response{Error: err.Error()})
			return
		}
		jsonResponse(w, Response{Result: result})
	})

	// start http server over port 8000 and log for fatal errors.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Response struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func jsonResponse(w http.ResponseWriter, x interface{}) {
	// encode x as json
	bytes, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	// Write to ResponseWriter and send to client
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func whoisQuery(data string) (string, error) {
	// Run whoisquery.
	response, err := whois.Fetch(data)
	if err != nil {
		return "", err
	}
	return string(response.Body), nil
}
