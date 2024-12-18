package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) getTotal(w http.ResponseWriter, r *http.Request) {

	debut, err := s.svcs.CRUD().GetTotal(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(debut); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
