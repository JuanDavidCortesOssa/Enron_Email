package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Source struct {
	ID                      int    `json:"ID"`
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	XCC                     string `json:"X-cc"`
	XBCC                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Message                 string `json:"Message"`
}

type fullResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Documents []struct {
			Source Source `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/emails", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Term string `json:"term"`
			From int    `json:"from"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var result []byte
		if requestBody.Term == "" {
			result = ZincSearchEmptyTermPost(requestBody.From)
		} else {
			result = ZincSearchPost(requestBody.Term, requestBody.From)
		}

		var fullResponse fullResponse
		err := json.Unmarshal(result, &fullResponse)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			http.Error(w, "Error parsing JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(fullResponse); err != nil {
			log.Println("Error encoding JSON response:", err)
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}
	})

	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Term string `json:"term"`
			From int    `json:"from"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(requestBody); err != nil {
			log.Println("Error encoding JSON response:", err)
			http.Error(w, "Error generating response", http.StatusInternalServerError)
		}
	})

	r.Get("/emails", func(w http.ResponseWriter, r *http.Request) {
		term := r.URL.Query().Get("term")
		fromStr := r.URL.Query().Get("from")

		from, err := strconv.Atoi(fromStr)
		if err != nil || from < 0 {
			http.Error(w, "'from' must be a non-negative integer", http.StatusBadRequest)
			return
		}

		result := ZincSearchPost(term, from)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", r)
}

func ZincSearchPost(term string, from int) []byte {
	query := fmt.Sprintf(`{
        "search_type": "match",
        "query": {
            "term": "%s"
        },
        "from": %d,
        "max_results": 20,
		"sort_fields": ["-Date"]
    }`, term, from)

	req, err := http.NewRequest("POST", "http://zincsearch:4080/api/Enron_Data/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func ZincSearchEmptyTermPost(from int) []byte {
	query := fmt.Sprintf(`{
        "search_type": "alldocuments",
        "from": %d,
        "max_results": 20,
		"sort_fields": ["-Date"]
    }`, from)

	req, err := http.NewRequest("POST", "http://zincsearch:4080/api/Enron_Data/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	return body
}
