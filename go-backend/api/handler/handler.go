package handler

import (
	"github.com/Porcari-Karim/GeoMeeting/api/handler/auth"
	"net/http"
)

var GlobalHandler *http.ServeMux

func Init() {
	GlobalHandler = http.NewServeMux()
	InitBase(GlobalHandler)
	auth.InitAuth(GlobalHandler)
}
