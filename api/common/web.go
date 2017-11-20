package common

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, model interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		logger.Println("JSONResponse: err:", err)
	}
}
