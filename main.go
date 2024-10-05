package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getTitle(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("URL fetching failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status encountered: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	title := doc.Find("title").Text()
	return title, nil
}

func worker(id int, urls <-chan string, results chan<- string) {
	for url := range urls {
		fmt.Printf("worker %d is fetching URL: %s\n", id, url)
		res, err := getTitle(url)
		if err != nil {
			results <- fmt.Sprintf("ERR while fetching %s: %v", url, err)
		} else {
			results <- fmt.Sprintf("Title of %s: %s", url, res)
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://invalid-url.com",
	}
	urlChan := make(chan string, len(urls))
	resChan := make(chan string, len(urls))
	const numWokers = 3

	for i := 1; i <= numWokers; i++ {
		go worker(i, urlChan, resChan)
	}

	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-resChan)
	}
}
