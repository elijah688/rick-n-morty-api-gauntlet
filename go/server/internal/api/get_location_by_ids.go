package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) getLocationByIDs(w http.ResponseWriter, r *http.Request) {

	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %s", err), http.StatusBadRequest)
		return
	}

	characters, err := s.svcs.CRUD().GetLocationByID(r.Context(), reqBody.IDs)
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
