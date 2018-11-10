package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	log.Println(message)
	Json(w, code, map[string]interface{}{
		"code":    code,
		"message": message,
	})
}

func Json(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
