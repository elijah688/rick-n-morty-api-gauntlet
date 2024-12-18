package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *Server) getCharacterByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID parameter", http.StatusBadRequest)
		return
	}

	char, err := s.svcs.Gateway().CompileCharacter(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if char == nil {
		http.Error(w, "character not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(char); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
