package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func chirpHandler(w http.ResponseWriter, r *http.Request){

	bad_words := []string{"kerfuffle", "sharbert", "fornax"}

	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Cleaned string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	sentence := params.Body

	words := strings.Split(sentence, " ")


	for i := 0; i < len(words); i++ {
		for j := 0; j < len(bad_words); j++ {
			if strings.ToLower(words[i]) == bad_words[j] {
				words[i] = "****"

			}
		}
	}

	cleaned_sentence := strings.Join(words, " ")

	respBody := returnVals {
		Cleaned: cleaned_sentence,
	}
	
	respondWithJSON(w, 200, respBody)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 500, "Something went wrong")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type errorVals struct{
		Error string `json:"error"`
	}

	respBody := errorVals {
		Error: msg,
	}

	respondWithJSON(w, code, respBody)
}