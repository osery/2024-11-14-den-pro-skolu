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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func listMessages(w http.ResponseWriter, r *http.Request) {
	for _, m := range messages {
		fmt.Fprintf(w, "%s\n", m)
	}
}
