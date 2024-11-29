package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/avicienna99/reserverings_app_fonteyn.git/app/api"
	"github.com/avicienna99/reserverings_app_fonteyn.git/app/db"
)

func main() {

	config, err := db.LoadConfig("../config/db_config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbConn, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/houses", api.GetHousesHandler(dbConn, config.Table))

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
