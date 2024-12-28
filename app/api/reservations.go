package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Reservation struct {
	HouseID   int    `json:"house_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func ReservationHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var reservation Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			log.Printf("Failed to decode reservation data: %v", err)
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		query := `INSERT INTO reservations (house_id, name, email, start_date, end_date) VALUES (?, ?, ?, ?, ?)`
		_, err := dbConn.Exec(query, reservation.HouseID, reservation.Name, reservation.Email, reservation.StartDate, reservation.EndDate)
		if err != nil {
			log.Printf("Failed to insert reservation: %v", err)
			http.Error(w, "Failed to save reservation", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
