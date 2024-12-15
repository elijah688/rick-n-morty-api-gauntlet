package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) getCharacters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 0 {
		http.Error(w, "invalid limit parameter", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil || offset < 0 {
		http.Error(w, "invalid offset parameter", http.StatusBadRequest)
		return
	}

	characters, err := s.svcs.CRUD().GetCharacters(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
