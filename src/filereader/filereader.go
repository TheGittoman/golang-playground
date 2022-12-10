package filereader

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// LoadFile load words from the file
func LoadFile(dataChan *chan string, fileName string) (*os.File, int64) {
	// opening the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Opening the file: ", err)
	}
	// scanner for reading the file
	stats, err := file.Stat()
	if err != nil {
		log.Fatalln("Error Opening the file: ", err)
	}
	return file, stats.Size()
}

// ScanFile scans the file for data and closes the data channel
func ScanFile(file *os.File, dataChan *chan string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*dataChan <- scanner.Text()
	}
	close(*dataChan)
}

// SaveWords read words from the data channel
func SaveWords(wg *sync.WaitGroup, dataChan *chan string, concurrency int) []string {
	wg.Add(1)
	words := []string{}
	// read work concurrently
	go func() {
		go addWords(wg, dataChan, &words, &concurrency)
	}()
	wg.Wait()
	return words
}

// addWords adds all the data to a list from a channel
func addWords(wg *sync.WaitGroup, dataChan *chan string, words *[]string, concurrency *int) {
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		for word := range *dataChan {
			*words = append(*words, word)
		}
		wg.Done()
	}
	wg.Done()
}
