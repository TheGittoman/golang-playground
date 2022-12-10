package main

import (
	"fmt"
	"playground/src/filereader"
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

	// load data in to a channel
	file, size := filereader.LoadFile(&dataChan, fileName)
	defer file.Close()
	fmt.Println("size of file: ", size)

	go filereader.ScanFile(file, &dataChan)

	// send data to a channel
	words := filereader.SaveWords(wg, &dataChan, concurrency)

	fmt.Println("time elapsed: ", time.Since(start))
	fmt.Println("length of wordlist: ", len(words))
	fmt.Println("first word: ", words[0], "last word: ", words[len(words)-1])
}
