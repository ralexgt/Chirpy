package main

import (
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpRetrieveById(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirp")
		return
	}
	id_value := r.PathValue("chirpId")
	id, err := strconv.Atoi(id_value)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get id path value")
		return
	}
	chirp := Chirp{}
	for _, dbChirp := range dbChirps {
		if dbChirp.ID == id {
			chirp = Chirp{
				ID:   dbChirp.ID,
				Body: dbChirp.Body,
			}
		}
	}
	if (chirp == Chirp{}) {
		respondWithError(w, http.StatusNotFound, "No chirp found")
	} else {
		respondWithJSON(w, http.StatusOK, chirp)
	}
}