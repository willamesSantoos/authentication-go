package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/willamesSantoos/authentication/internal/repository"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := repository.SelectAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}
