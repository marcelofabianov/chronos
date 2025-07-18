package v1

import (
	"encoding/json"
	"net/http"
)

// @TODO: Finalizar func de retornar users

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "users retrieved successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
