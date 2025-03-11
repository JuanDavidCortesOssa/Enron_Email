package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

const (
	maxFolderReaders = 5
	maxFileReaders   = 10
	maxBatchSize     = 50
	databasePath     = "C:/Users/JUAN DAVID/Desktop/TechnicalTestData/enron_mail_20110402/maildir/"
)

var (
	wg                    sync.WaitGroup
	folderReaderSemaphore chan struct{}
	fileReadSemaphore     chan struct{}
	emailChannel          chan string
)

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// runtime.SetBlockProfileRate(1)

	start := time.Now()
	log.Println("Process started...")

	emailChannel = make(chan string, 100)

	go batchProcessor()

	fillBatchAsync(databasePath)

	wg.Wait()
	close(emailChannel)

	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Data sent. Total time: %s\n", elapsed)
}

func batchProcessor() {
	batchSizeLimit := 1048576

	var batch []string
	var currentBatchSize int

	for emailData := range emailChannel {

		if len(emailData) > batchSizeLimit {
			log.Printf("Skipping email: Size=%d exceeds maxEmailSize=%d", len(emailData), batchSizeLimit)

			ZincSearchDataIndexer(emailData)

			continue
		}

		indexEntry := `{"index": {"_index": "Enron_Data"}}`
		entrySize := len(indexEntry) + len(emailData) + 2

		if currentBatchSize+entrySize > batchSizeLimit {
			ZincSearchBulkIndexer(batch)
			log.Printf("Batch sent: Size=%d bytes, Emails=%d\n", currentBatchSize, len(batch)/2)
			batch = batch[:0]
			currentBatchSize = 0
		}

		batch = append(batch, indexEntry, emailData)
		currentBatchSize += entrySize
	}

	log.Println("Bulking data 1 ", len(batch))

	if len(batch) > 0 {
		ZincSearchBulkIndexer(batch)
		log.Println("Bulking data ", len(batch))
		//log.Printf("Final batch sent: Size=%d bytes, Emails=%d\n", currentBatchSize, len(batch)/2)
	}
}

func fillBatchAsync(basePath string) {
	folderReaderSemaphore = make(chan struct{}, maxFolderReaders)
	fileReadSemaphore = make(chan struct{}, maxFileReaders)

	user_list := list_all_folders(basePath)
	for _, user := range user_list {
		var userPath string = basePath + user
		wg.Add(1)
		folderReaderSemaphore <- struct{}{}
		go func(userPath string) {
			defer wg.Done()
			defer func() { <-folderReaderSemaphore }()
			traverseFolders(userPath)
		}(userPath)
	}

}

func traverseFolders(basePath string) {
	entries, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		entryPath := basePath + "/" + entry.Name()
		if entry.IsDir() {
			traverseFolders(entryPath)
		} else {
			wg.Add(1)
			fileReadSemaphore <- struct{}{}
			go func(path string) {
				defer wg.Done()
				defer func() { <-fileReadSemaphore }()

				emailData := EmailParser(path)
				emailChannel <- emailData
			}(entryPath)
		}
	}
}

func list_all_folders(folder_name string) []string {
	files, err := os.ReadDir(folder_name)
	if err != nil {
		log.Fatal(err)
	}
	var list_folders []string
	for _, f := range files {
		list_folders = append(list_folders, f.Name())
	}
	return list_folders
}
