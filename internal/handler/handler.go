package handler

import (
	"encoding/json"
	"net/http"

	"hello-go/internal/greeting"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	uppercase := r.URL.Query().Get("uppercase") == "true"

	if name != "" {
		if err := greeting.ValidateName(name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}

	msg := greeting.GenerateGreeting(name)
	msg = greeting.FormatGreeting(msg, uppercase)

	_ = json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
