package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avicienna99/reserverings_app_fonteyn.git/app/db"
)

// GetHousesHandler handles HTTP requests to fetch house data
func GetHousesHandler(dbConn *sql.DB, tableName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Query the database
		query := fmt.Sprintf("SELECT id, name, description, price, availability FROM %s", tableName)
		rows, err := dbConn.Query(query)
		if err != nil {
			log.Printf("Query failed: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Parse rows into a slice of House structs
		var houses []db.House
		for rows.Next() {
			var house db.House
			if err := rows.Scan(&house.ID, &house.Name, &house.Description, &house.Price, &house.Availability); err != nil {
				log.Printf("Row scan failed: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			houses = append(houses, house)
		}

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(houses)
	}
}
