package controller

import (
	"github.com/rupalbarman/2048-game-webserver/model"
	"os"
	"log"
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

	port := os.Getenv("PORT")
	//port:="8080"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", game)
	http.HandleFunc("/left", left)
	http.HandleFunc("/right", right)
	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)
	http.HandleFunc("/about", about)
	//http.Handle("/view", http.FileServer(http.Dir("../2048-game-webserver/view")))
	http.ListenAndServe(":" + port, nil)
}