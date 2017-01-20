package controller

import (
	"github.com/rupalbarman/model"
	
	"net/http"
)

func game(w http.ResponseWriter, r *http.Request) {
	if r.Method== "GET" {
		renderTemplate(w, "../src/view/game.html", b_copy)
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
		renderTemplate(w, "../src/view/about.html", b_copy)
	} else {
		r.ParseForm()
	}
}