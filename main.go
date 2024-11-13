package main

import (
	"html/template"
	"net/http"
)

var messages []string

func main() {
	http.HandleFunc("GET /", listMessages)
	http.HandleFunc("POST /send", newMessage)
	http.ListenAndServe(":8080", nil)
}

func listMessages(w http.ResponseWriter, r *http.Request) {
	page, _ := template.ParseGlob("html/messages.gohtml")
	page.Execute(w, messages)
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	messages = append(r.Form["message"], messages...)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
