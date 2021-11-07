package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avocadohooman/Golang-Rest-Api-Todo/domain"
)

func (s *Server) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := domain.RegisterPayload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			response := map[string]string{"error": err.Error()}
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		fmt.Println(payload)

		// user, err := s.domain.Register()
	}
}
