package main

import (
	"log"
	"net/http"
	"os"

	"reserverings_app_fonteyn/app/api"
	"reserverings_app_fonteyn/app/db"
)

// test 123
func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	http.Handle("/", http.FileServer(http.Dir("./app/static")))
	http.HandleFunc("/api/houses", api.GetHousesHandler(dbConn, os.Getenv("DB_TABLE")))
	http.HandleFunc("/api/reservations", api.ReservationHandler(dbConn))

	log.Fatal(http.ListenAndServe(":80", nil))
}
