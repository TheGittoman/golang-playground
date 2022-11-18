package multithreading

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// MultiThreading test project
func MultiThreading() {
	start := time.Now()
	file, err := os.Open(".\\data\\Hitchhiker's Guide to Galaxy.txt")
	if err != nil {
		panic(err)
	}
	// fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	jobs := make(chan string)
	results := make(chan string)
	defer file.Close()
	wg := new(sync.WaitGroup)

	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go waitNumbers(jobs, results, wg)
	}

	count := 0
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			count++
			jobs <- (fmt.Sprint(count) + scanner.Text())
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	// writeFile, err := os.Create("./data/combined.dat")
	// if err != nil {
	// 	panic(err)
	// }

	// writer := bufio.NewWriter(writeFile)
	for i := range results {
		i = strings.ToLower(i)
		if strings.Contains(i, "marvin") {
			fmt.Println("Found")
			break
		}
	}

	elapsed := time.Since(start)
	log.Println("Took: ", elapsed)
}

func waitNumbers(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range jobs {
		results <- i
	}
}

func printFunc(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Println("Waiting...  ", id)
		time.Sleep(1 * time.Second)
	}
}

// MultiThreading2 short wg test
func MultiThreading2() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printFunc(&wg, i)
	}
	wg.Wait()
}
