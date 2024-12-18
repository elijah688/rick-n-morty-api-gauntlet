package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) getDebuts(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request body: %s", err), http.StatusBadRequest)
		return
	}

	debut, err := s.svcs.CRUD().GetDebutByIDs(r.Context(), reqBody.IDs)
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
