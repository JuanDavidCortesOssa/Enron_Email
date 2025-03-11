package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type email struct {
	ID                        int       `json:"ID"`
	Message_ID                string    `json:"Message-ID"`
	Date                      time.Time `json:"Date"`
	From                      string    `json:"from"`
	To                        string    `json:"to"`
	Subject                   string    `json:"subject"`
	Mime_Version              string    `json:"Mime-Version"`
	Content_Type              string    `json:"Content-Type"`
	Content_Transfer_Encoding string    `json:"Content-Transfer-Encoding"`
	X_From                    string    `json:"X-From"`
	X_To                      string    `json:"X-To"`
	X_cc                      string    `json:"X-cc"`
	X_bcc                     string    `json:""X-bcc"`
	X_Folder                  string    `json:"X-Folder"`
	X_Origin                  string    `json:"X-Origin"`
	X_FileName                string    `json:"X-FileName"`
	Message                   string    `json:"Message"`
}

func EmailParser(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	var emailData email
	reader := bufio.NewReader(file)
	bodyStarted := false
	var bodyBuilder strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading file %s: %v", filePath, err)
			}
			break // End of file
		}

		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			// Empty line indicates the start of the body
			bodyStarted = true
			continue
		}

		if bodyStarted {
			// Collect body lines
			bodyBuilder.WriteString(line + "\n")
		} else {
			// Parse headers
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) != 2 {
				continue
			}
			key, value := parts[0], parts[1]

			switch key {
			case "Message-ID":
				emailData.Message_ID = value
			case "Date":
				emailData.Date = parseDateType(value)
			case "From":
				emailData.From = value
			case "To":
				emailData.To = value
			case "Subject":
				emailData.Subject = value
			case "Mime-Version":
				emailData.Mime_Version = value
			case "Content-Type":
				emailData.Content_Type = value
			case "Content-Transfer-Encoding":
				emailData.Content_Transfer_Encoding = value
			case "X-From":
				emailData.X_From = value
			case "X-To":
				emailData.X_To = value
			case "X-cc":
				emailData.X_cc = value
			case "X-bcc":
				emailData.X_bcc = value
			case "X-Folder":
				emailData.X_Folder = value
			case "X-Origin":
				emailData.X_Origin = value
			case "X-FileName":
				emailData.X_FileName = value
			}
		}
	}

	emailData.Message = bodyBuilder.String()

	// Convert emailData to JSON
	jsonData, err := json.Marshal(emailData)
	if err != nil {
		log.Printf("Error marshalling to JSON for file %s: %v", filePath, err)
		return ""
	}

	//log.Println(jsonData)

	return string(jsonData)
}

func parseDateType(date string) time.Time {

	format := "Mon, _2 Jan 2006 15:04:05 -0700 (MST)"
	parseDateType, err := time.Parse(format, date)

	if err != nil {
		log.Println("Error parsing date ", date)
		return time.Now()
	}

	return parseDateType
}
