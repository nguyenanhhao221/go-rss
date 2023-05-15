package main

import "net/http"

// handlerReadiness check if the server is ready
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{}) //struct{}{} when output as json should be an empty json object
}
