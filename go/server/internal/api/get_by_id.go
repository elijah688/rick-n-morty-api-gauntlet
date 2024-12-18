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

	item, err := s.svcs.CRUD().GetCharacterByID(r.Context(), id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, fmt.Sprintf("no item with id %d found", id), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
