package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/igor-sn/user-register/src/models"
	"log"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error converting id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil && rows == 0 {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("WARN: %v users have been deleted", rows)
	}

	w.WriteHeader(http.StatusOK)
	return
}
