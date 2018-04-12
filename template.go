package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/define", define)
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./assets/img"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./assets/css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./assets/js"))))
	//this hides the asset part of the location... now can reference img/file.yyy
	log.Fatal(http.ListenAndServe(":4001", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", "Home")
}
func define(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "define.gohtml", "Define")
}
func about(w http.ResponseWriter, r *http.Request) {
	type customData struct {
		Title   string
		Members []string
	}
	cd := customData{
		Title:   "About",
		Members: []string{"Moneypenny", "Bond", "Q", "M"},
	}
	tpl.ExecuteTemplate(w, "about.gohtml", cd)
}
