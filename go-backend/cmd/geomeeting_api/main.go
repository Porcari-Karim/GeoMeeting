package main

import (
	"github.com/Porcari-Karim/GeoMeeting/internal/storage/cache"
	"github.com/Porcari-Karim/GeoMeeting/internal/storage/db"
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

	db.Connect()
	defer func() {
		err := db.DB.Close()
		if err != nil {
			panic(err)
		}
	}()

	cache.Connect()
	defer func() {
		err := cache.Cache.Close()
		if err != nil {
			panic(err)
		}
	}()

	handler.Init()
	err = http.ListenAndServe(":8080", handler.GlobalHandler)
	if err != nil {
		panic(err)
	}

}
