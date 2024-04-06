package handler

import (
	"html/template"
	"net/http"
)

var baseHandler *http.ServeMux
var indexTemplate *template.Template

func initBase(rootMux *http.ServeMux) {
	indexTemplate = template.Must(template.ParseFiles("./web/templates/index.html"))

	baseHandler = http.NewServeMux()
	rootMux.Handle("/", baseHandler)
	staticFilesHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))
	baseHandler.Handle("GET /static/", staticFilesHandler)

	baseHandler.HandleFunc("/", indexHandler)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
