package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/willamesSantoos/authentication/internal/models"
	"github.com/willamesSantoos/authentication/internal/repository"
	"github.com/willamesSantoos/authentication/internal/util"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalln("Error when decoding json: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(user.Email) == 0 || len(user.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide email and password to obtain the token"))
		return
	}

	isExist, err := repository.SelectUser(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Oops, error when trying to find user"))
	}

	if isExist {
		token, err := util.GetToken(user.Email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", token)
			w.WriteHeader(http.StatusOK)

			response := map[string]string {
				"token": token,
			}

			jsonResponse, err := json.Marshal(response)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(jsonResponse)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}

		result, err := util.VerifyToken(tokenString)

		if err != nil || !result {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	})
}
