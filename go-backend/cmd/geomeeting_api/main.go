package main

import (
	"log"
	"net/http"

	"github.com/Porcari-Karim/GeoMeeting/api/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	handler.Init()
	http.ListenAndServe(":8080", handler.GlobalHandler)

}
