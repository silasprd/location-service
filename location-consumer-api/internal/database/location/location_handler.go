package location

import (
	"encoding/json"
	"net/http"

	database "github.com/silasprd/sailor-location-service/location-consumer-api/internal/database/connection"
)

func FindAllHandler(w http.ResponseWriter, r *http.Request) {
	locationDB := Location{DB: database.DB}
	locations, err := locationDB.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
