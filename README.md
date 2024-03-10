#### Getting started

clone the repo
```bash
git clone https://github.com/prkagrawal/task-of-go-server
```

move inside the project directory
```bash
cd task-of-go-server
```

Run the server
```bash
go run main.go
```

Open a browser and navigate to http://localhost:8080
You should see the json response, the server is up


#### API USAGE

Endpoint: /keyword

Request
```json
{
  "keyword": "TATAPOWER"
}
```
Response
```json
{
  "original_keyword": "TATA Power",
  "new_keywords": [],
  "matched_chunks": [
    "Tata Power, formerly a part of the three entities jointly known as Tata Electric Companies, is one of India's largest Integrated Power Company.",
    "Tata Power, together with its subsidiaries & joint entities, has a generation capacity of 14,464 MW of which 39% comes from clean energy sources. The company has the distinction of being among the top private players in each sector of the value chain including solar rooftop and value-added services.",
  ],
  "summary": "Summary generated from llm"
}
```

#### Testing
To test the server with concurrent requests
Use Apache Benchmark (ab) tool, run
```bash
ab -n 20000 -c 5000 http://localhost:8080
```
This will send 20000 requests with 5000 concurrent requests, you should see the server handling the requests concurrently
