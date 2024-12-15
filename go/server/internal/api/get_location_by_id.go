package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *Server) getLocationByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID parameter", http.StatusBadRequest)
		return
	}

	characters, err := s.svcs.CRUD().GetLocationByID(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
