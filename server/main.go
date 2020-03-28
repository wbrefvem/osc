package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Crawlable represents a crawable unit, i.e. domain
type Crawlable struct {
	URL string
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing GET...")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var c Crawlable

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Crawling domain %s\n", c.URL)

	cmd := exec.Command(
		"scrapy",
		"crawl",
		"osc",
		"-a",
		fmt.Sprintf("allowed_domains=%s", c.URL),
		"-a",
		fmt.Sprintf("start_urls=%s", c.URL),
	)

	cmd.Dir = "/Users/wrefvem/go/src/github.com/wbrefvem/osc-gateway"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Printf("ERROR: %s", err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	log.Println("handling request...")
	if r.Method == http.MethodGet {
		handleGet(w, r)
	} else if r.Method == http.MethodPost {
		handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	log.Println("Listening for crawl requests...")

	mux := http.NewServeMux()
	mux.Handle("/crawl", http.HandlerFunc(handleRequests))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
