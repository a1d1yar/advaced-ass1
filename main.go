package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Message string `json:"message"`
}

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

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if reqBody.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Println("Received message:", reqBody.Message)

	resBody := ResponseBody{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(resBody)
}
