package handler

import (
	"html/template"
	"net/http"
)

var authHandler *http.ServeMux
var loginTemplate, logoutTemplate *template.Template

func initAuth(rootMux *http.ServeMux) {
	loginTemplate = template.Must(template.ParseFiles("./web/templates/login.html"))
	logoutTemplate = template.Must(template.ParseFiles("./web/templates/logout.html"))

	authHandler = http.NewServeMux()
	rootMux.Handle("/auth/", http.StripPrefix("/auth", authHandler))

	authHandler.HandleFunc("GET /login", loginHandler)
	authHandler.HandleFunc("POST /login", loginHandler)
	authHandler.HandleFunc("GET /logout", logoutHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := loginTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	err := logoutTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
