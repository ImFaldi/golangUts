package main

import (
	"Personal"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/name", Personal.PersonalHandler)
	mux.HandleFunc("/allData", Personal.AllDataHandler)

	http.ListenAndServe(":5050", mux)
}
