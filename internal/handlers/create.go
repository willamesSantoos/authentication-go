package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/willamesSantoos/authentication/internal/models"
	"github.com/willamesSantoos/authentication/internal/repository"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Contente-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalln("Error when decoding json: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := repository.Insert(user)

	var resp models.ServerResponse

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		resp = models.ServerResponse{
			Message: fmt.Sprintf("Oops! Something went wrong, Err: %v", err),
		}
	} else {
		w.WriteHeader(http.StatusCreated)

		resp = models.ServerResponse{
			Message: fmt.Sprintf("Registration created successfully! Id: %d", id),
		}
	}

	json.NewEncoder(w).Encode(resp)
}
