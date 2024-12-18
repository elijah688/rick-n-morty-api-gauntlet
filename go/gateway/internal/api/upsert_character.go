package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/model"
)

func (s *Server) upsertCharacter(w http.ResponseWriter, r *http.Request) {
	var character model.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := character.Validate(); err != nil {
		http.Error(w, "invalid character", http.StatusBadRequest)
		return
	}
	res, err := s.svcs.Gateway().UpsertCharacter(r.Context(), character)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
