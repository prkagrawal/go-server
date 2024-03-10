package main

// simple http server to handle concurrent requests, and return a json response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "time"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Handling request")
	response := Response{Message: "Hello, World!"}
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

	// make ap api call to localhost:8081/keyword with POST request
	// and send the keyword var in the request body
	// and get the response

	response, err := http.Post("http://localhost:8081/keyword", "application/json", {"keyword": keyword})
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse the response body and store it in a json object
	var res map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&res)
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


// Run the server
// go run main.go
// Open a browser and navigate to http://localhost:8080
// You should see the json response

// To test the server with concurrent requests
// Use Apache Benchmark (ab) tool
// ab -n 1000 -c 100 http://localhost:8080
// This will send 1000 requests with 100 concurrent requests
// You should see the server handling the requests concurrently
// and returning the json response
