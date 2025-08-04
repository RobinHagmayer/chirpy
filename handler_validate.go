package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type requestParameters struct {
		Body string `json:"body"`
	}

	type response struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	reqParams := requestParameters{}
	err := decoder.Decode(&reqParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(reqParams.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
		return
	}

	respondWithJson(w, http.StatusOK, response{
		Valid: true,
	})
}
