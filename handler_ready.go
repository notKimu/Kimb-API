package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "API up and running -w-")
}
