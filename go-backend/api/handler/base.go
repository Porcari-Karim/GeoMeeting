package handler

import (
	"html/template"
	"net/http"
)

var indexTemplate *template.Template

func InitBase(rootMux *http.ServeMux) {
	indexTemplate = template.Must(template.ParseFiles("./web/templates/index.html"))

	staticFilesHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))
	rootMux.Handle("GET /static/", staticFilesHandler)
	rootMux.HandleFunc("GET /", indexHandler)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
