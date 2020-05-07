package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	output, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer output.Close()

	if err != nil {
		log.Fatal(err)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		if _, err = output.WriteString(<-ch); err != nil {
			log.Fatal(err)
		}
	}

	_, err = fmt.Fprintf(output, "%.2fs elapsed\n", time.Since(start).Seconds())
	if err != nil {
		log.Fatal(err)
	}

	if err = output.Close(); err != nil {
		log.Fatal(err)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
