package handler

import (
	"net/http"
)

var GlobalHandler *http.ServeMux

func Init() {
	GlobalHandler = http.NewServeMux()
	initBase(GlobalHandler)
	initAuth(GlobalHandler)
}
