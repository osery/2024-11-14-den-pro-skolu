package main

import (
	"html/template"
	"net/http"
)

var messages []string

func main() {
	http.HandleFunc("GET /", listMessages)
	http.HandleFunc("POST /send", newMessage)
	err := http.ListenAndServeTLS(":443", "certs/fullchain.pem", "certs/privkey.pem", nil)
	if err != nil {
		panic(err)
	}
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
