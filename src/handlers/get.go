package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/igor-sn/user-register/src/models"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error get user: %v", err)
		http.Error(w, "Error get user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
