package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var concurrency int = 50
var fileName string = "./data/English_words_470k.txt"

// tree for word list that finds the words as fast as possible
func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	dataChan := make(chan string)
	go loadFile(&dataChan, fileName)
	words := saveWords(wg, &dataChan, concurrency)

	// words := []string{}
	// for word := range dataChan {
	// 	words = append(words, word)
	// }

	fmt.Println("length of wordlist: ", len(words))
	fmt.Println("first word: ", words[0], "last word: ", words[len(words)-1])
	fmt.Println("time elapsed: ", time.Since(start))
}

// load words from the file
func loadFile(dataChan *chan string, fileName string) {
	// opening the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Opening the file: ", err)
	}
	defer file.Close()
	// scanner for reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*dataChan <- scanner.Text()
	}
	close(*dataChan)
}

// read words from the data channel
func saveWords(wg *sync.WaitGroup, dataChan *chan string, concurrency int) []string {
	wg.Add(1)
	words := []string{}
	// read work concurrently
	go func() {
		for i := 0; i < concurrency; i++ {
			wg.Add(1)
			for word := range *dataChan {
				words = append(words, word)
			}
			wg.Done()
		}
		wg.Done()
	}()
	wg.Wait()
	return words
}
