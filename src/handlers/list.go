package handlers

import (
	"encoding/json"
	"github.com/igor-sn/user-register/src/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAll()
	if err != nil {
		log.Printf("Error listing users: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
