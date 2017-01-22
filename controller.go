package main

import (
	"github.com/rupalbarman/2048-game-webserver/model"
	
	"os"
	"log"
	"net/http"
	"html/template"
)

var b_copy *model.Board

func game(w http.ResponseWriter, r *http.Request) {
	if r.Method== "GET" {
		renderTemplate(w, "view/game.html", b_copy)
	} else {
		r.ParseForm()
	}
}

func left(w http.ResponseWriter, r *http.Request) {
	//LEFT
	model.SwipeHorizontal(b_copy, true)
	model.GenerateTile(b_copy, 2, len(b_copy.Matrix), len(b_copy.Matrix[0]))
	http.Redirect(w, r, "/", http.StatusFound)
}

func right(w http.ResponseWriter, r *http.Request) {
	//RIGHt
	model.SwipeHorizontal(b_copy, false)
	model.GenerateTile(b_copy, 2, len(b_copy.Matrix), len(b_copy.Matrix[0]))
	http.Redirect(w, r, "/", http.StatusFound)
}

func up(w http.ResponseWriter, r *http.Request) {
	//UP
	model.SwipeVertical(b_copy, true)
	model.GenerateTile(b_copy, 2, len(b_copy.Matrix), len(b_copy.Matrix[0]))
	http.Redirect(w, r, "/", http.StatusFound)
}

func down(w http.ResponseWriter, r *http.Request) {
	//DOWN
	model.SwipeVertical(b_copy, false)
	model.GenerateTile(b_copy, 2, len(b_copy.Matrix), len(b_copy.Matrix[0]))
	http.Redirect(w, r, "/", http.StatusFound)
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method== "GET" {
		renderTemplate(w, "view/about.html", b_copy)
	} else {
		r.ParseForm()
	}
}

func renderTemplate(w http.ResponseWriter, templ string, b *model.Board) {
	t, _ := template.ParseFiles(templ)
	t.Execute(w, b)
}

func run(b *model.Board) {
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