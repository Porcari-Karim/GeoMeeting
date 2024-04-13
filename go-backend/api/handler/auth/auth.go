package auth

import (
	"html/template"
	"net/http"
)

var loginTemplate, logoutTemplate *template.Template

func InitAuth(rootMux *http.ServeMux) {
	loginTemplate = template.Must(template.ParseFiles("./web/templates/login.html"))
	logoutTemplate = template.Must(template.ParseFiles("./web/templates/logout.html"))

	rootMux.HandleFunc("GET /auth/login", loginHandler)
	rootMux.HandleFunc("POST /auth/login", loginHandler)
	rootMux.HandleFunc("GET /auth/logout", logoutHandler)

	rootMux.HandleFunc("GET /auth/o/google/", googleOAuthHandler)
	rootMux.HandleFunc("GET /auth/o/google/callback/", googleOAuthCallbackHandler)
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
