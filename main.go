package main

import (
	"encoding/json"
	// "fmt"
	"bytes"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Handling request")
	response := Response{Message: "Hi, the api is live"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println("Sending response")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Handling request")
	response := Response{Message: "Ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func handleKeywordResponse(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Handling request")

	// parse the request body and store it in a map
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if the request body is empty, return an error
	if len(request) == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	// check if the request body contains a keyword
	keyword, ok := request["keyword"].(string)
	if !ok {
		http.Error(w, "Request body does not contain a keyword", http.StatusBadRequest)
		return
	}

	// make an POST req to localhost:8081/keyword
	// and send the keyword in the request body
	url := "http://localhost:8081/keyword"

	// JSON body
	body := []byte(`{
		"keyword": "` + keyword + `"
	}`)

	// make the POST request
	response, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse the response body and store it in a json object
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if the response body is empty, return an error
	if len(res) == 0 {
		http.Error(w, "Response body is empty", http.StatusInternalServerError)
		return
	}

	// return the response from the api call
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/keyword", handleKeywordResponse)
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
