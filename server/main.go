package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// Crawlable represents a crawable unit, i.e. domain
type Crawlable struct {
	URL string
}

func processURL(rawURL string) (*url.URL, error) {
	log.Printf("process URL %s\n", rawURL)
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, fmt.Errorf("URL scheme is invalid")
	}

	return parsed, nil
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing GET...")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing POST")
	var c Crawlable

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		http.Error(w, "JSON body is malformed", http.StatusBadRequest)
		return
	}

	parsedURL, err := processURL(c.URL)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	domainAndPort := strings.Split(parsedURL.Host, ":")
	bareDomain := domainAndPort[0]

	log.Printf("Crawling domain %s\n", bareDomain)

	cmd := exec.Command(
		"scrapy",
		"crawl",
		"osc",
		"-a",
		fmt.Sprintf("allowed_domains=%s,", bareDomain),
		"-a",
		fmt.Sprintf("start_urls=%s,", fmt.Sprintf("%s://%s/%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path)),
	)

	cmd.Dir = os.Getenv("WORK_DIR")
	if cmd.Dir == "" {
		cmd.Dir = "/opt/crawler"
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		http.Error(w, "failed to start crawl command", http.StatusInternalServerError)
	}

	io.WriteString(w, fmt.Sprintf("started crawl for domain %s", c.URL))
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
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
