package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

// amount of goprocesses to use for reading the file later
var concurrency int = 100

// file to read the data from
var fileName string = "./data/English_words_470k.txt"

func main() {
	// open file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Error Opening a file: ", err)
	}
	// cloese file at function exit
	defer file.Close()

	// waitgroup to control when to exit and close channels
	wg := new(sync.WaitGroup)
	// string channel for work
	dataChan := make(chan string)

	// get some work
	go func() {
		// new scanner for getting the lines from the file
		inRead := bufio.NewScanner(file)
		// scan the file and send lines to the string channel
		for inRead.Scan() {
			dataChan <- inRead.Text()
		}
		// close the channel after we are done with reading the file
		// to avoid deadlock
		close(dataChan)
	}()

	// list of strings for saving the contents of the file
	wordlist := []string{}
	// reading the data concurrently in the word list
	for i := 0; i < concurrency; i++ {
		// adding worgroup to control the exit of the program
		wg.Add(1)
		go func() {
			for v := range dataChan {
				wordlist = append(wordlist, v)
			}
			// we are done with the thread so let the workgroup know
			wg.Done()
		}()
		// wait for all the threads to be done
		wg.Wait()
	}
	// write the info about the file
	fmt.Println("Lines in file: ",
		len(wordlist), " last word: ",
		wordlist[len(wordlist)-1])
}
