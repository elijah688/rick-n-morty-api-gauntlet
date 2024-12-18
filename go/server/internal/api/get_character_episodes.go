package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) getCharacterEpisodesByIDs(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %s", err), http.StatusBadRequest)
		return
	}

	episodes, err := s.svcs.CRUD().GetCharacterEpisodes(r.Context(), reqBody.IDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(episodes); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
