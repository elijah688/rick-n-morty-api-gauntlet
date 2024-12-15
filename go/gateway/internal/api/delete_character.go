package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *Server) deleteCharacter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID parameter", http.StatusBadRequest)
		return
	}

	if err := s.svcs.Gateway().DeleteCharacter(r.Context(), id); err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, id)))

}
