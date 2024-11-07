package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed html/messages.gohtml
var html embed.FS
var tmpl = template.Must(template.ParseFS(html, "html/messages.gohtml"))

var messages []string

func main() {
	http.HandleFunc("GET /", listMessages)
	http.HandleFunc("POST /send", newMessage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func listMessages(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, messages)
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	messages = append(r.Form["message"], messages...)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
