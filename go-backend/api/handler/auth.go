package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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

	authHandler.HandleFunc("/o/google/", googleOAuthHandler)
	authHandler.HandleFunc("/o/google/callback/", googleOAuthCallbackHandler)
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

func googleOAuthHandler(w http.ResponseWriter, r *http.Request) {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	redirectURI := "http://" + r.Host + "/auth/o/google/callback/"
	scope := "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile"
	authRedirectURL := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?response_type=code&scope=%s&redirect_uri=%s&client_id=%s", scope, redirectURI, clientId)
	http.Redirect(w, r, authRedirectURL, http.StatusTemporaryRedirect)
}

func googleOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	userInfo, err := getGoogleUserInfoFromCallback(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userInfo.Email))
}
