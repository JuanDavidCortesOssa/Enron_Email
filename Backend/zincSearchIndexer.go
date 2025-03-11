package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 10,
	},
	Timeout: time.Second * 10,
}

func ZincSearchBulkIndexer(data []string) {
	if len(data) == 0 {
		return
	}
	bulkData := strings.Join(data, "\n") + "\n"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:4080/api/_bulk", strings.NewReader(bulkData))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	//log.Print(bulkData)

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/x-ndjson")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return
	}
	log.Printf("Response: %s, body: %s", resp.Status, string(body))
}

func ZincSearchDataIndexer(data string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:4080/api/Enron_Data/_doc", strings.NewReader(data))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = client.Do(req)
		if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
			break
		}
		if resp != nil {
			resp.Body.Close()
		}
		log.Println("Retrying request...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Println("Error sending request after retries:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("Failed to index document: %s, data length: %d", resp.Status, len(data))
	}
}
