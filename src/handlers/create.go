package handlers

import (
	"encoding/json"
	"github.com/igor-sn/user-register/src/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding user data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	res := map[string]int64{"id": id}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
