// Run the server
// go run main.go
// Open a browser and navigate to http://localhost:8080
// You should see the json response

// To test the server with concurrent requests
// Use Apache Benchmark (ab) tool
// ab -n 20000 -c 5000 http://localhost:8080
// This will send 20000 requests with 5000 concurrent requests
// You should see the server handling the requests concurrently
// and returning the json response