package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type result struct {
	url      string
	duration time.Duration
	err      error
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	ch := make(chan result)

	for _, url := range urls {
		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)
			if err == nil {
				resp.Body.Close()
			}
			ch <- result{url: u, duration: time.Since(start), err: err}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-ch
		if r.err != nil {
			fmt.Printf("%s -> error: %v\n", r.url, r.err)
		} else {
			fmt.Printf("%s -> %v\n", r.url, r.duration)
		}
	}
}
