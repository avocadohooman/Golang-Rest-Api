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
			badRequestResponse(w, err)
			return
		}
		fmt.Println("payload", payload)
		// user, err := s.domain.Register()
	}
}
