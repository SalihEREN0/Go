package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.youtube.com/",
		"https://www.google.com",
		"https://go.dev",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer resp.Body.Close()

			fmt.Println(url, resp.Status)
		}(url)
	}

	wg.Wait()

	fmt.Println("Done")
}
