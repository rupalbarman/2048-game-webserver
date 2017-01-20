package controller

import (
	"github.com/rupalbarman/model"
	
	"net/http"
	"html/template"
)

var b_copy *model.Board

func renderTemplate(w http.ResponseWriter, templ string, b *model.Board) {
	t, _ := template.ParseFiles(templ)
	t.Execute(w, b)
}

func Run(b *model.Board) {
	// b_copy points to the original board
	b_copy= b

	http.HandleFunc("/", game)
	http.HandleFunc("/left", left)
	http.HandleFunc("/right", right)
	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)
	http.HandleFunc("/about", about)
	//http.Handle("/view", http.FileServer(http.Dir("../src/view")))
	http.ListenAndServe(":8000", nil)
}