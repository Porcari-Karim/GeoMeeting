package handler

import (
	"github.com/Porcari-Karim/GeoMeeting/internal/storage/db"
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

	var queryResult struct {
		mail string
		name string
	}

	err = db.DB.QueryRow("select * from m_user").Scan(&queryResult.mail, &queryResult.name)
	if err != nil {
		panic(err)
	}

	_, err = w.Write([]byte(queryResult.mail))
	if err != nil {
		panic(err)
	}

}
