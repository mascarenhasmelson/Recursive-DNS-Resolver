package main

import(
  "fmt"
  "net/http"
  "log"
  "time"
  "github.com/mascarenhasmelson/Recursive-DNS-Resolver/core"
)


func dnsQuery(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("domain")
	if query == "" {
		http.Error(w, "Missing 'domain' query parameter", http.StatusBadRequest)
		return
	}
  // Set headers for streaming response
  	w.Header().Set("Content-Type", "text/plain")
  	w.Header().Set("Transfer-Encoding", "chunked")
  	flusher, ok := w.(http.Flusher)
  	if !ok {
  		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
  		return
  	}

  	// Start DNS query and stream responses
  	fmt.Fprintf(w, "Starting DNS Query for: %s\n", query)
  	flusher.Flush()

  	// Call the recursive DNS resolver function
  	answers, err := core.Returnresults(query, w, flusher)

  	if err != nil {
  		fmt.Fprintf(w, "Error: %v\n", err)
  		flusher.Flush()
  		return
  	}

  	// Stream each result
    fmt.Fprintln(w, "Final Resolved Records:")
  	for _, record := range answers {
  		fmt.Fprintf(w, "Received Answer: %s\n", record.String())
  		flusher.Flush()
  		time.Sleep(time.Millisecond * 500) // Simulating delay per request
  	}

  	fmt.Fprintf(w, "Query Completed.\n")
  	flusher.Flush()


}
func main(){

  http.HandleFunc("/query", dnsQuery)

  	fmt.Println("Starting API server on :8080")
  	log.Fatal(http.ListenAndServe(":8080", nil))


}
