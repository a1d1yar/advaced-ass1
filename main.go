package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RequestBody represents the structure of the incoming JSON data
type RequestBody struct {
	Message string `json:"message"`
}

// ResponseBody represents the structure of the outgoing JSON response
type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON from the request body
	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Check if the "message" field is present
	if reqBody.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	// Print the message to the server console
	fmt.Println("Received message:", reqBody.Message)

	// Respond with success message
	resBody := ResponseBody{
		Status:  "success",
		Message: "Data successfully received",
	}

	// Encode and send the JSON response
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(resBody)
}
