package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func sleep(t time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Start operation: %s", t.String())
	time.Sleep(t)
	log.Printf("Finish operation: %s", t.String())
}

func main() {

	var wg sync.WaitGroup

	datafile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer datafile.Close()

	logfile, err := os.Create("result.txt")
	if err != nil {
		os.Exit(1)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	scanner := bufio.NewScanner(datafile)
	for scanner.Scan() {
		t, _ := time.ParseDuration(scanner.Text())
		wg.Add(1)
		go sleep(t, &wg)
	}

	wg.Wait()
	// t, _ := time.ParseDuration("7h3m45s")
	// fmt.Printf("%.0f s\n", t.Seconds())
}
