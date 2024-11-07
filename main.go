package main

import (
	"fmt"
	"net/http"
)

var messages = []string{
	"zprava 1",
	"zprava 2",
}

func main() {
	http.HandleFunc("GET /", listMessages)
	http.ListenAndServe(":8080", nil)
}

func listMessages(w http.ResponseWriter, r *http.Request) {
	for _, m := range messages {
		fmt.Fprintf(w, "%s\n", m)
	}
}
