package main

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver/utils"
)


func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case utils.Error:
		log.Printf("HTTP %d - %s ", e.Status(), e)
		respondWithJSON(w, e.Status(), e.Error(),nil)
	default:
		respondWithJSON(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError),nil)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{},headers map[string]string) {
	resp, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	for header := range headers {
		w.Header().Set(header,headers[header])
	}

	w.WriteHeader(code)
	w.Write(resp)
}
