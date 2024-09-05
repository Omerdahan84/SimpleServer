package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	// we have a bug
	if code > 499{
		log.Println("Respond with 5xx erorr :", msg)
	}
	// create a new struct for return JSON reesponse
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}


// respondWithJSON is a helper function to send a JSON response back to the client.
// It takes the following parameters:
// - w: the http.ResponseWriter used to send the response.
// - code: the HTTP status code (e.g., 200 for OK, 404 for Not Found).
// - payload: the data to be sent as a JSON object in the response body.
func respondWithJson( w http.ResponseWriter, code int, payload interface{})  {
	// Convert the payload (data) into a JSON-formatted byte array.

	data, err := json.Marshal(payload)
	// If an error occurs during the conversion to JSON, log the error
	// and respond with a 500 Internal Server Error status code.
	if err!=nil{
		log.Println("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return 
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) // indicates everything went well 
	w.Write(data)

}