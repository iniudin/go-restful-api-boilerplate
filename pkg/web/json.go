package web

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")

	payload, err := json.Marshal(data)
	if err != nil {
		http.Error(w, `{"status":"failed", "error":"failed to encode response"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(payload)
}

func ParseJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
