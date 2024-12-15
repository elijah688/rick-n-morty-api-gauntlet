package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *Server) getDebut(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID parameter", http.StatusBadRequest)
		return
	}

	debut, err := s.svcs.CRUD().GetDebutByID(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Println(debut)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(debut); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
