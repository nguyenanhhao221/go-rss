package main

import "net/http"

// handlerReadiness check if the server is ready
func handlerError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something went wrong") //struct{}{} when output as json should be an empty json object
}
