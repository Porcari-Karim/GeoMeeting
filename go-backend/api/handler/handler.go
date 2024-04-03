package handler

import (
	"net/http"
	"os"
)

var GlobalHandler = http.NewServeMux()

func Init() {
	GlobalHandler.HandleFunc("GET /", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	jwtKey := os.Getenv("JWT_KEY")
	_, err := w.Write([]byte(jwtKey))
	if err != nil {
		return
	}
}
