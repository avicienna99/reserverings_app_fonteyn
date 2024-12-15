package main

import (
	"log"
	"net/http"

	"reserverings_app_fonteyn/app/api"
	"reserverings_app_fonteyn/app/db"
)

// test
func main() {

	config, err := db.LoadConfig("./config/db_config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbConn, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	http.Handle("/", http.FileServer(http.Dir("./app/static")))

	http.HandleFunc("/api/houses", api.GetHousesHandler(dbConn, config.Table))

	log.Fatal(http.ListenAndServe(":80", nil))
}
